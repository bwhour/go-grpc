/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary server is an example server.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bwhour/go-grpc/lib/grpc"
	"github.com/bwhour/go-grpc/lib/grpc/codes"
	"github.com/bwhour/go-grpc/lib/grpc/credentials/alts"
	ecpb "github.com/bwhour/go-grpc/lib/grpc/examples/features/proto/echo"
	"github.com/bwhour/go-grpc/lib/grpc/status"
)

var port = flag.Int("port", 50051, "the port to serve on")

type ecServer struct{}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
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
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create alts based credential.
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())

	s := grpc.NewServer(grpc.Creds(altsTC))

	// Register EchoServer on the server.
	ecpb.RegisterEchoServer(s, &ecServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
