package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bwhour/go-grpc/lib/go-log"
	"github.com/bwhour/go-grpc/lib/grpc"
	"github.com/bwhour/go-grpc/lib/grpc/codes"
	"github.com/bwhour/go-grpc/lib/grpc/status"
	ecpb "github.com/bwhour/go-grpc/lib/grpc/test/examples/features/proto/echo"
	"github.com/bwhour/go-grpc/server"
)

type ecServer struct {
	addr string
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}
func (s *ecServer) ServerStreamingEcho(*ecpb.EchoRequest, ecpb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) ClientStreamingEcho(ecpb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) BidirectionalStreamingEcho(ecpb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

func main() {
	log.Init(&log.Config{
		Output:       "std",
		Level:        "DEBUG",
		ShowFileLine: true,
	})
	defer log.Close()

	srv := server.New(&server.Config{
		Name:          "test",
		Address:       ":51234",
		SocketTimeout: 1 * time.Second,
		TCPCheck:      false,
	})

	gs := grpc.NewServer()
	ecpb.RegisterEchoServer(gs, &ecServer{addr: "localhost:51234"})

	err := srv.Serve(gs)
	if err != nil {
		log.Fatalf("server err :%v", err)
	}
}
