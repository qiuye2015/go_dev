package main

import (
	"SecKill/Proxy/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	etcd_client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)
var (
	redisPool *redis.Pool
	etcdClient *etcd_client.Client
)

func initSec() (err error) {
	err = initLog()
	if err != nil{
		logs.Error("init logger failed,err:%v",err)
	}
	err = initRedis()
	if err !=nil{
		logs.Error("init redis failed,err:",err)
		return
	}
	err = initEtcd()
	if err !=nil {
		logs.Error("init etcd failed,err:",err)
		return
	}

	err = loadSecConf()
	if err !=nil{
		logs.Error("loadSecConf failed,err:",err)
		return
	}

	initSecProductWatcher()

	service.InitService(secKillConf)

	logs.Info("init sec succ")

	return
}

func initLog() (err error){
	convertLogLevel := func(level string) int {
		switch (level) {
		case "debug":
			return logs.LevelDebug
		case "warn":
			return logs.LevelWarn
		case "info":
			return logs.LevelInfo
		case "trace":
			return logs.LevelTrace
		}
		return logs.LevelDebug
	}
	config :=make(map[string]interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)
	configStr,err:=json.Marshal(config)
	if err !=nil{
		fmt.Println("json Marshal failed,err:",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	return
}

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisConf.RedisIdleTimeout)*time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisConf.RedisAddr)
		},
	}
	conn :=redisPool.Get()
	defer conn.Close()
	_,err =conn.Do("ping")
	if err !=nil{
		logs.Error("ping redis error",err)
		return
	}

	return
}

func initEtcd() (err error) {

	cli,err:= etcd_client.New(etcd_client.Config{
		Endpoints: []string{secKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(secKillConf.EtcdConf.Timeout)*time.Second,
	})
	if err !=nil{
		logs.Error("connect etcd failed,err:",err)
		return
	}
	etcdClient=cli
	logs.Info("init etct succ")
	return
}

func loadSecConf() (err error){
	resp,err:=etcdClient.Get(context.Background(), secKillConf.EtcdConf.EtcdSecProductKey)
	if err !=nil{
		logs.Error("get [%s] from etcd faild err:%v", secKillConf.EtcdConf.EtcdSecProductKey,err)
		return
	}
	var secProductInfo []service.SecProductInfo
	for k,v:=range resp.Kvs{
		logs.Debug("key[%v],value[%v]",k,v)
		err = json.Unmarshal(v.Value,&secProductInfo)
		if err !=nil{
			logs.Error("json unmarshal sec product info faied,err:%v",v)
			return
		}
		logs.Debug("sec info conf is [%v]",secProductInfo)
	}
	updateSecProductInfo(secProductInfo)
	return
}

func initSecProductWatcher(){
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
	return
}

func watchSecProductKey(key string){
	logs.Debug("begin watch key:%s",key)
	for{
		rch:=etcdClient.Watch(context.Background(),key)
		var secProductInfo []service.SecProductInfo
		getConfSucc := true

		for wresp :=range rch{
			for _,ev:=range wresp.Events{
				if ev.Type == mvccpb.DELETE{
					logs.Warn("Key[%s]'s config deleted",key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key)==key{
					err := json.Unmarshal(ev.Kv.Value,&secProductInfo)
					if err !=nil{
						logs.Error("Key[%s],Unmarshal[%s],err:%v",key,err)
						getConfSucc =false
						continue
					}
				}
				logs.Debug("get config from etcd,%s %q : %q",ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ,%v",secProductInfo)
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}

func updateSecProductInfo(secProductInfo []service.SecProductInfo) {
	var tmp map[int]*service.SecProductInfo =make(map[int]*service.SecProductInfo,1024)
	for _,v:=range secProductInfo{
		tmp[v.ProductId] =&v
	}
	secKillConf.RWSecProductLock.Lock()
	secKillConf.SecProductInfoMap = tmp
	secKillConf.RWSecProductLock.Unlock()
}
