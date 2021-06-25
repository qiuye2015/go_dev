package main

import (
	//"RpcCode/param"
	"fmt"
	"github.com/qiuye2015/go_dev/param"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	//var req float32
	//req = 3
	////同步方式调用
	//var resp *float32
	//err = client.Call("MathUtil.CalulatorCircalArea",req,&resp)
	//if err != nil{
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	////异步调用
	//var resSync *float32
	//syncCall :=client.Go("MathUtil.CalulatorCircalArea",req,&resSync,nil)
	//replayDone := <- syncCall.Done
	//fmt.Println(replayDone)
	//fmt.Println(*resSync)

	var result *float32
	addParam := &param.AddParam{Args1: 1, Args2: 2}
	err = client.Call("fjp.Add", addParam, &result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*result)
}
