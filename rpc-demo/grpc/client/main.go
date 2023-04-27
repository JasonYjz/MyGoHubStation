package main

import (
	"context"
	"go-hub-station/rpc-demo/grpc/appmon"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50500", grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to Dial. err:%s", err.Error())
	}

	defer conn.Close()

	c := appmon.NewGreeterClient(conn)
	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(ctx, &appmon.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greeting: %s", r.Message)
}
