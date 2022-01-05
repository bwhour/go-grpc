package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/bwhour/go-grpc/lib/grpc"
	log "github.com/bwhour/go-log"
)

var (
	privateBlocks []*net.IPNet
)

func init() {
	for _, b := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "100.64.0.0/10"} {
		if _, block, err := net.ParseCIDR(b); err == nil {
			privateBlocks = append(privateBlocks, block)
		}
	}
}

const (
	defaultTCPCheck         = 5 * time.Second
	defaultTTL              = 10 * time.Second
	defaultRegisterInternal = 5 * time.Second
)

// Server 是 grpc server 封装
type Server struct {
	config *Config

	done   chan struct{}
	close  chan struct{}
	closed int32

	wg  sync.WaitGroup
	ttl time.Duration
}

// New 创建新的 grpc server。
func New(config *Config) *Server {
	if config.SocketTimeout <= 0 {
		config.SocketTimeout = DefaultSocketTimeout
	}

	if config.MetricsName == "" {
		config.MetricsName = DefaultMetricsName
	}

	if config.Address == "" {
		log.Fatalf("server address is empty")
	}

	as := strings.Split(config.Address, ":")

	config.addr = as[0]
	config.port, _ = strconv.Atoi(as[1])

	return &Server{
		config: config,
		done:   make(chan struct{}),
		close:  make(chan struct{}),
		ttl:    defaultTTL,
	}
}

// Serve 启动 grpc server。
// 这个函数会一直阻塞直到服务被停止，会自动监听 SIGINT 和 SIGTERM。
func (s *Server) Serve(gs *grpc.Server) (err error) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.port))
	if err != nil {
		log.Fatalf("grpc: fail to listen %v", err)
	}
	var addr string
	if s.config.addr != "" {
		addr = s.config.addr
	}

	addr, err = extract()
	if err != nil {
		return
	}

	// 捕获各种信号。
	c := make(chan os.Signal, 1)
	defer close(c)

	exit := make(chan error)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	defer signal.Stop(c)

	go func() {
		select {
		case <-s.done:
			gs.Stop()
			exit <- nil
		case sig := <-c:
			log.Errorf("_sensego_grpc_server||msg=SIGINT or SIGTERM is notified.||sig=%v", sig)
			gs.Stop()
			exit <- fmt.Errorf("grpc server stops.||signal=%v", sig)
		}

		if state := atomic.SwapInt32(&s.closed, 2); state == 1 || state == 0 {
			close(s.close)
		}
	}()

	log.Infof("_sensego_grpc_server||msg=grpc server is ready to start.||addr=%v", fmt.Sprintf("%s:%d", addr, s.config.port))

	if err := gs.Serve(lis); err != nil {
		log.Errorf("fail to start grpc server.||err=%v||addr=%v", err, fmt.Sprintf("%s:%d", addr, s.config.port))
		return err
	}

	return <-exit
}

// Close 用来停止 server。
// 实现了 gracefully shutdown，会等所有请求结束才退出。
func (s *Server) Close() error {
	if s.closed == 2 {
		return nil
	}

	if atomic.SwapInt32(&s.closed, 1) != 0 {
		return nil
	}

	close(s.done)
	select {
	case <-s.close:
	}

	s.wg.Wait()
	return nil
}

// extract returns a real ip
func extract() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("Failed to get interface addresses! Err: %v", err)
	}

	var ipAddr []byte
	var publicIP []byte

	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if ip.To4() == nil {
			continue
		}

		if !isPrivateIP(ip.String()) {
			publicIP = ip
			continue
		}

		ipAddr = ip
		break
	}

	// return private ip
	if ipAddr != nil {
		return net.IP(ipAddr).String(), nil
	}

	// return public or virtual ip
	if publicIP != nil {
		return net.IP(publicIP).String(), nil
	}

	return "", fmt.Errorf("No IP address found, and explicit IP not provided")
}

func isPrivateIP(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}
