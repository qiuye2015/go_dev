package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type SKProductInfo struct {
	ProductID int
	Total     int // 总量
	Stock     int // 库存
	StartTime int
	EndTime   int
	Status    int
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.22.5.25:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		panic(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//resp, err := cli.Get(ctx, "greeting")
	//resp, err := cli.Put(ctx, "sample_key", "sample_value")
	var prod []SKProductInfo
	prod = append(prod, SKProductInfo{
		ProductID: 1000,
		Total:     1024,
		Stock:     1024,
		StartTime: 1621558461,
		EndTime:   1621568461,
		Status:    0,
	})

	prod = append(prod, SKProductInfo{
		ProductID: 2000,
		Total:     1024,
		Stock:     1024,
		StartTime: 1621558461,
		EndTime:   1621568461,
		Status:    0,
	})
	prodStr, err := json.Marshal(prod)
	if err != nil {
		fmt.Println("json Marshal failed,err:", err)
		return
	}
	key := "/fjp/seckill/product"
	resp, err := cli.Put(ctx, key, string(prodStr))
	defer cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	fmt.Println(resp)
	resp2, err := cli.Get(ctx, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp2)
}
