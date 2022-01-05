# go-grpc
go-grpc for k8s client

基于 [grpc-go]( https://github.com/grpc/grpc-go ) 封装了 grpc 包，基础版本为v1.23.0，实现了基于kubernetes的client端服务发现。


提供了 client 和 server。封装了 服务发现，重试，metrics(TODO)的逻辑。使调用者无需关心

> 需要配合 github.com/bwhour/go-grpc/lib/protobuf/protoc-gen-sensego 这个生成器使用


## example

```go
// client 的使用
client := NewWebSrvClient(&client.Config{
  Name: "webservice.sensego", // service-name 或 service-name.namespace
  Retry: 3, // 失败后重试次数
  SocketTimeout: 3 * time.Second, // 超时时间
})

```
