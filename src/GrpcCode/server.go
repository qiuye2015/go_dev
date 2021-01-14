package main

import (
	"GrpcCode/message"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

type OrderServiceImpl struct {
}
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error){
	orderMap :=map[string]message.OrderInfo{
		"1":{
			OrderId: "1",
			OrderName:"衣服",
			OrderStatus: "已付款",
		},
		"2":{
			OrderId:     "2",
			OrderName:   "零食",
			OrderStatus: "已付款",
		},
		"3":{
			OrderId:     "3",
			OrderName:   "饮料",
			OrderStatus: "未付款",
		},
	}
	var response *message.OrderInfo
	current:=time.Now().Unix()
	if (request.TimeStamp>current){
		*response= message.OrderInfo{
			OrderId:     "0",
			OrderName:   "",
			OrderStatus: "订单异常",
		}
	}else{
		result:=orderMap[request.OrderId]
		if result.OrderId!=""{
			fmt.Println(result)
			return &result,nil
		}else{
			return nil, errors.New("server error fjp")
		}
	}
	return response,nil
}

func main() {
	server:=grpc.NewServer()
	message.RegisterOrderServiceServer(server,new(OrderServiceImpl))
	lis,err:=net.Listen("tcp","localhost:8090")
	if err !=nil{
		panic(err.Error())
	}
	server.Serve(lis)

}
