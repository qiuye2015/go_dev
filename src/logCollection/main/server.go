package main

import (
	//"kafka"
	//"tailf"
	kafka "logCollection/kafka"
	tailf "logCollection/tailf"
	"time"

	"github.com/beego/beego/core/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed,err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	//fmt.Printf("read msg:%s,topic:%s\n", msg.Msg, msg.Topic)
	kafka.SendTokafka(msg.Msg, msg.Topic)
	return
}
