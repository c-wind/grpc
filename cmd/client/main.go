package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"

	"github.com/c-wind/grpc/meta"
)

var serverAddr = flag.String("server", "0.0.0.0:10000", "echo server addr")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ctx := context.Background()
	req := &meta.HelloRequest{
		Name: "tttttttttttttt",
	}

	echo := meta.NewGreeterClient(conn)
	resp, err := echo.SayHello(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp:%s\n", resp.GetMessage())

}
