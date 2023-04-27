package main

import (
	"context"
	"go-hub-station/rpc-demo/grpc/appmon"
	"google.golang.org/grpc"
	"log"
	"net"
)

type My struct {
}

func (m *My) SayHello(ctx context.Context, request *appmon.HelloRequest) (*appmon.HelloReply, error) {
	log.Printf("server say hello")
	return &appmon.HelloReply{Message: "hello"}, nil
}

func main() {
	server := grpc.NewServer()
	appmon.RegisterGreeterServer(server, new(My))

	listener, err := net.Listen("tcp", ":50500")
	if err != nil {
		log.Printf("failed to Listen. err:%s", err.Error())
	}

	if err := server.Serve(listener); err != nil {
		log.Printf("failed to Serve. err:%s", err.Error())
		return
	}
}
