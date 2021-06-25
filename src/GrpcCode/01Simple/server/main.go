package main

import (
	"context"
	proto "github.com/qiuye2015/go_dev/GrpcCode/01Simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type LoginServer_fjp struct{}

//判断用户名，密码是否为root,123456，验证正确即返回
func (l LoginServer_fjp) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginRes, error) {
	if req.Username == "root" && req.Password == "123456" {
		return &proto.LoginRes{Msg: "true"}, nil
	} else {
		return &proto.LoginRes{Msg: "false"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	// 构建一个新的服务端对象
	s := grpc.NewServer()
	// 向这个服务端对象注册服务
	proto.RegisterLoginServer(s, &LoginServer_fjp{})
	// 注册服务端反射服务
	reflection.Register(s)
	// 启动服务
	s.Serve(lis)

	// 可配合ctx实现服务端的动态终止
	//s.Stop()
}
