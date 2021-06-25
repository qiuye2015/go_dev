package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/qiuye2015/go_dev/crawler_distributed/config"
	"github.com/qiuye2015/go_dev/crawler_distributed/persist"
	"github.com/qiuye2015/go_dev/crawler_distributed/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, `the port for me to listen on (eg: --port ":1234")`)

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
	//log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	fmt.Println("starting serverPRC...")
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
		elastic.SetURL("http://10.22.5.25:9200/"),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServerRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
