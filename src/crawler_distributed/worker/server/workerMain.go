package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, `the port for me to listen on (eg: --port ":9000")`)

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	//log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
