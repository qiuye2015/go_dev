package main

import (
	"flag"
	"github.com/qiuye2015/go_dev/crawler/engine"
	"github.com/qiuye2015/go_dev/crawler/scheduler"
	citylist "github.com/qiuye2015/go_dev/crawler/zhenai/parser"
	"github.com/qiuye2015/go_dev/crawler_distributed/config"
	itemsaver "github.com/qiuye2015/go_dev/crawler_distributed/persist/client"
	"github.com/qiuye2015/go_dev/crawler_distributed/rpcsupport"
	worker "github.com/qiuye2015/go_dev/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("woker_hosts", "", `worker hosts (comma separated) (eg --woker_hosts ":9000,：9001")`)
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	//itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e3 := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e3.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(citylist.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
