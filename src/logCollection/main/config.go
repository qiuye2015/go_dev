package main

import (
	"errors"
	"fmt"

	//"tailf"
	tailf "logCollection/tailf"

	//"github.com/astaxie/beego/config"
	"github.com/beego/beego/core/config"
)

type Config struct {
	logLevel     string
	logPath      string
	kafkaAddr    string
	chanSize     int
	collectConfs []tailf.CollectConf

	etcdAddr string
}

var (
	appConfig *Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath, _ = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic, _ = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}
	appConfig.collectConfs = append(appConfig.collectConfs, cc)
	return
}
func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed,err:", err)
		return
	}
	appConfig = &Config{}
	appConfig.logLevel, _ = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath, _ = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	appConfig.kafkaAddr, _ = conf.String("kafka::kafka_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	appConfig.etcdAddr, _ = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collectConf err:%v", err)
		return
	}
	return
}
