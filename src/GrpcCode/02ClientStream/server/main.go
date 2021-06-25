package main

import (
	proto "github.com/qiuye2015/go_dev/GrpcCode/02ClientStream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
)

type UploadServer_fjp struct {
}

func (*UploadServer_fjp) Upload(server proto.Upload_UploadServer) error {
	for {
		// 循环接收客户端传递的流数据
		recv, err := server.Recv()

		//检测到EOF(客户端调用close)
		if err == io.EOF {
			//发送res
			err := server.SendAndClose(&proto.UploadRes{Msg: "finish"})
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			return err
		}
		log.Printf("get a upload data package~ offset:%v, size:%v\n", recv.Offset, recv.Size)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":6012")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//构建一个新的服务端对象
	s := grpc.NewServer()
	//向这个服务端对象注册服务
	proto.RegisterUploadServer(s, &UploadServer_fjp{})
	//注册服务端反射服务
	reflection.Register(s)

	//启动服务
	s.Serve(lis)

	//可配合ctx实现服务端的动态终止
	//s.Stop()
}
