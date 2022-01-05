package server

import "time"

const (
	// DefaultSocketTimeout 是默认的 socket 超时时间。
	DefaultSocketTimeout = 5 * time.Second

	// DefaultMetricsName 是默认的 metrics 统计指标名，一般来说应该设置成个性化的名字。
	DefaultMetricsName = "grpc-server"
)

// Config 表示 grpc server 的配置
type Config struct {
	Name          string        `toml:"name"`
	Address       string        `toml:"address"` // 服务监听的地址，ip:prot (ip可以为空)
	addr          string        `toml:"addr"`
	port          int           `toml:"port"`
	SocketTimeout time.Duration `toml:"socket_timeout"` // 服务 socket 超时设置，默认是 DefaultSocketTimeout。

	MetricsName     string `toml:"metrics_name"`     // metrics 统计时候用的指标名字，默认是 DefaultMetricsName。
	MetricsDisabled bool   `toml:"metrics_disabled"` // 是否完全不发送 metrics 统计，默认 false。
}
