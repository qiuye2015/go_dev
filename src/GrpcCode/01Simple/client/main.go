package main

import (
	proto "GrpcCode/01Simple"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// 创建grpc连接
	grpcConn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// 通过grpc连接创建一个客户端势力对象
	client := proto.NewLoginClient(grpcConn)
	// 设置ctx超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 通过client客户端对象，调用Login函数
	res, err := client.Login(ctx, &proto.LoginReq{
		Username: "root",
		Password: "1234567",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("the login anser is", res.Msg)
}
