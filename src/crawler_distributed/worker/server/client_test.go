package main

import (
	"github.com/qiuye2015/go_dev/crawler_distributed/config"
	"github.com/qiuye2015/go_dev/crawler_distributed/rpcsupport"
	"github.com/qiuye2015/go_dev/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/aba",
		Parser: worker.SerializedParser{
			Name: config.ParseCity,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
