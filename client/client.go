package client

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/bwhour/go-grpc/lib/grpc/codes"
	"github.com/bwhour/go-grpc/lib/grpc/resolver/kuberesolver"
	"github.com/bwhour/go-grpc/lib/grpc/status"

	log "github.com/bwhour/go-grpc/lib/go-log"
	"github.com/bwhour/go-grpc/lib/go-trace"
	"github.com/bwhour/go-grpc/lib/grpc"
)

const (
	defaultSchema = "kubernetes"
)

// 各种特殊错误。
var (
	ErrNoConnection      = fmt.Errorf("grpc: no connection")
	ErrRetryTooManyTimes = fmt.Errorf("grpc retried too many times")
	ErrNeedRetry         = fmt.Errorf("grpc: need retry")
)

const (
	dltagGrpcSuccess = "_com_grpc_success"
	dltagGrpcFailure = "_com_grpc_failure"
)

// Client 封装服务发现，负载均衡，重试，metrics 等逻辑，调用者无需关心细节
type Client struct {
	config      *Config
	metricsName string
	cc          *grpc.ClientConn
}

// New 返回新的 client，透明的进行服务发现和负载均衡，使用者无需关心细节
func New(config *Config) *Client {
	if config.Retry < 0 {
		config.Retry = 0
	}

	if config.SocketTimeout <= 0 {
		config.SocketTimeout = DefaultSocketTimeout
	}

	metricsName := config.MetricsName

	if metricsName == "" {
		metricsName = DefaultMetricsName
	}

	if config.Name == "" {
		log.Fatalf("grpc: client name is empty")
	}

	kuberesolver.RegisterInCluster()

	roundrobinConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s:", defaultSchema, config.Name),
		grpc.WithBalancerName("round_robin"), // This sets the initial balancing policy.
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("grpc: client dail failed: %v", err)
	}

	client := &Client{
		config:      config,
		metricsName: metricsName,
		cc:          roundrobinConn,
	}

	return client
}

// Action 代表一个真正的调用 grpc 服务的动作。
// 如果 Action 返回 ErrNeedRetry，框架会自动进行重试。
type Action func(cc *grpc.ClientConn) (err error)

// Do 执行一个 grpc rpc 调用并且自动封装重试逻辑。
// 调用者需要在 action 里面实际创建 grpc client 并且发起调用。
//
// 例子：
//     c := client.New(config) // c 可以提前作为全局变量创建好，所有方法都是多线程安全的。
//
//     var resp *fooservice.FooResponse
//     err := c.Do(ctx, func(cc *grpc.ClientConn) (err error) {
//         // `fooservice.NewFooServiceClientFactory` 是 grpc 自动生成的代码。
//         foo := fooservice.NewFooServiceClientFactory(cc)
//         resp, err = foo.Bar(/* params */)
//         return
//     })
//
//     if err != nil {
//         // 错误处理。
//     }
//     response := r.(fooservice.FooResponse)
//     // 使用 response。
func (c *Client) Do(ctx context.Context, action Action) error {
	var funcName string

	if pc, _, _, ok := runtime.Caller(1); ok {
		f := runtime.FuncForPC(pc)
		funcName = path.Base(f.Name())

		idx := strings.LastIndex(funcName, ".")

		if idx >= 0 {
			funcName = funcName[idx+1:]
		}
	}

	return c.DoWithName(ctx, funcName, action)
}

// DoWithName 执行一个 grpc rpc 调用并且自动封装重试逻辑，
// 这个方法跟 Do 的差别是手动指定了 callee 的名字，方便各种统计。
func (c *Client) DoWithName(ctx context.Context, method string, action Action) error {
	retry := c.config.Retry

	for i := 0; i <= retry; i++ {
		err := c.doOnce(ctx, method, i, action)

		if !needRetry(err) {
			return err
		}
	}

	return ErrRetryTooManyTimes
}

func (c *Client) doOnce(ctx context.Context, method string, retried int, action Action) (err error) {
	start := time.Now()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("grpc: panic recovered [panic:%v]", e)
		}

		procTime := time.Now().Sub(start)
		procTimeSeconds := (procTime / time.Microsecond * time.Microsecond).Seconds()

		dltag := dltagGrpcSuccess
		errmsg := ""

		if err != nil {
			dltag = dltagGrpcFailure
			errmsg = fmt.Sprintf("||errmsg=%v", err)
		}

		log.Infof("%v||%v||server_name=%v||function_name=%v||proc_time=%v||retried=%v%v",
			dltag, trace.ContextString(ctx), c.metricsName, method, procTimeSeconds, retried, errmsg)

		// TODO: 增加 metrics 打点
	}()

	err = ctx.Err()

	if err != nil {
		return
	}

	deadline, ok := ctx.Deadline()
	timeout := c.config.SocketTimeout

	if ok {
		if delta := deadline.Sub(start); delta < timeout {
			if delta <= 0 {
				err = context.DeadlineExceeded
				return
			}

			timeout = delta
		}
	}

	err = callAction(action, c.cc)
	return
}

func needRetry(err error) bool {
	if err == nil {
		return false
	}

	s, ok := status.FromError(err)
	if ok {
		if s.Code() == codes.Unavailable {
			return true
		}
	}

	if err == ErrNeedRetry {
		return true
	}

	// TODO: 增加错误校验，需要重试的错误进行重试
	return false
}

func callAction(action Action, cc *grpc.ClientConn) error {
	return action(cc)
}

// Close 关闭 client
func (c *Client) Close() error {
	return c.cc.Close()
}
