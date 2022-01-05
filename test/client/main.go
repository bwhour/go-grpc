package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bwhour/go-grpc/client"
	"github.com/bwhour/go-grpc/lib/grpc"
	ecpb "github.com/bwhour/go-grpc/lib/grpc/test/examples/features/proto/echo"
)

func main() {
	config := &client.Config{
		Name:  "test",
		Retry: 3,
	}

	cli := client.New(config)

	var err error
	in := &ecpb.EchoRequest{
		Message: "Hello world",
	}
	out := new(ecpb.EchoResponse)

	ctx := context.Background()

	ac := func(cc *grpc.ClientConn) (err error) {
		err = cc.Invoke(ctx, "/grpc.examples.echo.Echo/UnaryEcho", in, out)
		if err != nil {
			return err
		}
		return nil
	}

	err = cli.DoWithName(ctx, "UnaryEcho", ac)
	if err != nil {
		log.Fatalf("err : %v", err)
	}

	fmt.Println("out = ", out)
}
