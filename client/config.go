package client

import "time"

const (
	// DefaultSocketTimeout 是默认 socket 超时。
	DefaultSocketTimeout = 500 * time.Millisecond

	// DefaultMetricsName 是默认的 metrics 统计指标名，一般来说应该设置成个性化的名字。
	DefaultMetricsName = "grpc-client"

	// DefaultTCPCheck 默认的服务发现 tcp check 时间
	DefaultTCPCheck = 5 * time.Second
)

// Config grpc client 的相关设置
type Config struct {
	Name          string        `toml:"name"`           // 服务的名字，eg: web-service
	Retry         int           `toml:"retry"`          // 调用服务接口失败后，重试的次数，默认为 0
	SocketTimeout time.Duration `toml:"socket_timeout"` // 客户端 socket 超时时间，默认 DefaultSocketTimeout。

	// Not use
	MetricsName     string `toml:"metrics_name"`     // metrics 统计时候用的指标名字，默认是 DefaultMetricsName。
	MetricsDisabled bool   `toml:"metrics_disabled"` // 是否完全不发送 metrics 统计，默认 false。
}
