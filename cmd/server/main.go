package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/c-wind/grpc/meta"
)

var port = flag.Int("port", 10000, "The server port")

type echoServer struct {
}

func (s *echoServer) SayHello(_ context.Context, req *meta.HelloRequest) (*meta.HelloReply, error) {
	resp := meta.HelloReply{
		Message: fmt.Sprintf("recv:%s, time:%v", req.GetName(), time.Now()),
	}
	log.Printf("recv:%v\n", req)
	return &resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server := echoServer{}
	meta.RegisterGreeterServer(grpcServer, &server)
	grpcServer.Serve(lis)
}
