package main

import (
	"github.com/qiuye2015/go_dev/RpcCode/param"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

// 一
// 该方法向外暴露
//func (mu *MathUtil) CalulatorCircalArea(req float32, resp *float32) error {
//	*resp = math.Pi * req * req
//	return nil
//}

// 二
func (mu *MathUtil) Add(param param.AddParam, resp *float32) error {
	*resp = param.Args1 + param.Args2
	return nil
}

func main() {
	//1. 初始化指针数据类型
	mathUtil := new(MathUtil)

	//2. 调用net/rpc包的功能将服务对象进行注册
	//err := rpc.Register(mathUtil)
	err := rpc.RegisterName("fjp", mathUtil)
	if err != nil {
		panic(err.Error())
	}
	//3. 通过该函数把mathUtil中提供的服务注册到Http协议上，方便调用者可以利用http的方式进行数据传输
	rpc.HandleHTTP()

	//4. 在特定端口进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
