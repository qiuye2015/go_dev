package main

import (
	"crawler/engine"
	"crawler/model"
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url: "",
		Id:  "12345",
		Payload: model.Profile{
			Age:      44,
			Name:     "余生有你",
			Marriage: "离异",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
