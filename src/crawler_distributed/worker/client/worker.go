package client

import (
	"github.com/qiuye2015/go_dev/crawler/engine"
	"github.com/qiuye2015/go_dev/crawler_distributed/config"
	"github.com/qiuye2015/go_dev/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
