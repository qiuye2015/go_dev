package main

import (
	"SecKill/Proxy/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var (
	secKillConf = &service.SecSkillConf{
		SecProductInfoMap: make(map[int]*service.SecProductInfo,1024),
	}
)

func initConfig() (err error) {
	redisAddr :=beego.AppConfig.String("RedisAddr")
	logs.Debug("read config succ,redis addr:%v",redisAddr)

	redisMaxIdle,err:=beego.AppConfig.Int("RedisMaxIdle")
	if err !=nil{
		err = fmt.Errorf("init config failed,RedisMaxIdle err %v", err)
		return
	}

	redisMaxActive,err:=beego.AppConfig.Int("RedisMaxActive")
	if err !=nil{
		err = fmt.Errorf("init config failed,RedisMaxActive err %vr", err)
		return
	}

	redisIdleTimeout,err:=beego.AppConfig.Int("RedisIdleTimeout")
	if err !=nil{
		err = fmt.Errorf("init config failed,RedisIdleTimeout err %v", err)
		return
	}

	etcdAddr :=beego.AppConfig.String("EtcdAddr")

	etcdTimeout,err:=beego.AppConfig.Int("etcdTimeout")
	if err !=nil{
		err = fmt.Errorf("init config failed,etcdTimeout err: %v", err)
		return
	}

	etcdKeyPrefix :=beego.AppConfig.String("etcdSecKeyPrefix")
	etcdKey :=beego.AppConfig.String("etcdSecKeyPrefix")

	if len(redisAddr)==0 ||len(etcdAddr)==0 || len(etcdKeyPrefix)==0  || len(etcdKey)==0{
		err = fmt.Errorf("init config failed,redis[%s] or etcd[%s] or etcdKeyPrefix[%s] or etcdKey[%s] config is null", redisAddr, etcdAddr,etcdKeyPrefix,etcdKey)
		return
	}

	secKillConf.RedisConf.RedisAddr =redisAddr
	secKillConf.RedisConf.RedisMaxIdle =redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive =redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout =redisIdleTimeout

	secKillConf.EtcdConf.EtcdAddr =etcdAddr
	secKillConf.EtcdConf.Timeout = etcdTimeout
	if strings.HasSuffix(etcdKeyPrefix,"/") ==false{
		etcdKeyPrefix = etcdKeyPrefix + "/"
	}
	secKillConf.EtcdConf.EtcdSecKeyPreifx = etcdKeyPrefix
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s",etcdKeyPrefix,etcdKey)


	secKillConf.LogPath=beego.AppConfig.String("logPath")
	secKillConf.LogLevel=beego.AppConfig.String("LogLevel")

	logs.Debug("read config succ,etcd  addr:%v",etcdAddr)
	return
}
