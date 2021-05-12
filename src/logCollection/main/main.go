package main

import (
	"fmt"
	"kafka"
	"tailf"

	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "../conf/log.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("init log err:%v\n", err)
		return
	}

	logs.Debug("load conf succ. config:%v", appConfig)

	err = tailf.InitTail(appConfig.collectConfs, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed ,err:%v", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed ,err:%v", err)
		return
	}

	logs.Debug("initialize succ")

	// go func() {
	// 	count := 0
	// 	for {
	// 		time.Sleep(time.Second)
	// 		logs.Debug("test %d", count)
	// 		count++
	// 	}
	// }()
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("main exited")

}
