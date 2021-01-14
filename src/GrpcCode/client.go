package main

import (
	"GrpcCode/message"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn,err:=grpc.Dial("localhost:8090",grpc.WithInsecure())
	if err !=nil{
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient :=message.NewOrderServiceClient(conn)
	orderRequest:= &message.OrderRequest{
		OrderId: "4",
		TimeStamp:time.Now().Unix(),
	}
	orderInfo,err:=orderServiceClient.GetOrderInfo(context.Background(),orderRequest)
	if orderInfo !=nil{
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
