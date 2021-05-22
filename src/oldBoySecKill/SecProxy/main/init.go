package main

import (
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/astaxie/beego/adapter"
	"github.com/astaxie/beego/core/logs"
	"github.com/gomodule/redigo/redis"
	clientv3 "go.Etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/api/mvccpb"
	"oldBoySecKill/SecProxy/service"
	"strings"
	"time"
)

var (
	gSecKillConf = &service.SecKillConf{
		SKProdInfosMap: make(map[int]*service.SKProductInfo, 1024),
	}
	gRedisPool  *redis.Pool
	gEtcdClient *clientv3.Client
)

func initConfig() (err error) {
	gSecKillConf.CookieSecretKey = beego.AppConfig.String("dev::cookieSecretKey")
	//logs
	gSecKillConf.LogPath = beego.AppConfig.String("dev::logPath")
	gSecKillConf.LogLevel = beego.AppConfig.String("dev::logLevel")
	//redis
	redisAddr := beego.AppConfig.String("dev::redisAddr")
	redisMaxIdle, err := beego.AppConfig.Int("dev::redisMaxIdle")
	if err != nil {
		err = fmt.Errorf("init config Int failed, err:%v", err)
		return
	}
	redisMaxActive, err := beego.AppConfig.Int("dev::redisMaxActive")
	if err != nil {
		err = fmt.Errorf("init config Int failed, err:%v", err)
		return
	}
	redisIdleTimeout, err := beego.AppConfig.Int("dev::redisIdleTimeout")
	if err != nil {
		err = fmt.Errorf("init config Int failed, err:%v", err)
		return
	}
	gSecKillConf.RedisConf.RedisAddr = redisAddr
	gSecKillConf.RedisConf.RedisMaxIdle = redisMaxIdle
	gSecKillConf.RedisConf.RedisMaxActive = redisMaxActive
	gSecKillConf.RedisConf.RedisIdleTimeout = redisIdleTimeout

	// etcd
	etcdAddr := beego.AppConfig.String("dev::etcdAddr")
	gSecKillConf.EtcdConf.EtcdAddr = etcdAddr
	etcdTimeout, err := beego.AppConfig.Int("dev::etcdTimeout")
	if err != nil {
		err = fmt.Errorf("init config Int failed, err:%v", err)
		return
	}
	etcdKeyPrefix := beego.AppConfig.String("dev::etcdKeyPrefix")
	if !strings.HasSuffix(etcdKeyPrefix, "/") {
		etcdKeyPrefix = etcdKeyPrefix + "/"
	}
	etcdProductKey := beego.AppConfig.String("dev::etcdProductKey")

	gSecKillConf.EtcdConf.EtcdTimeout = etcdTimeout
	gSecKillConf.EtcdConf.EtcdKeyPrefix = etcdKeyPrefix
	gSecKillConf.EtcdConf.EtcdProductKey = fmt.Sprintf("%s%s", gSecKillConf.EtcdConf.EtcdKeyPrefix, etcdProductKey)

	logs.Debug("redis addr:%v, etcd addr:%v", redisAddr, etcdAddr)
	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}
	logs.Info("initConfig success gSecKillConf = %v", gSecKillConf)
	return
}

func initSecKill() (err error) {
	if err = initLogs(); err != nil {
		logs.Error("init logger faild, err:%v", err)
		return
	}
	if err = initRedis(); err != nil {
		logs.Error("init redis faild, err:%v", err)
		return
	}
	if err = initEtcd(); err != nil {
		logs.Error("init etcd faild, err:%v", err)
		return
	}
	if err = loadSKProdInfos(); err != nil {
		logs.Error("loadSKProdInfos faild, err:%v", err)
		return
	}

	logs.Info("init seckill succ")
	return
}

func initLogs() (err error) {
	logConfig := make(map[string]interface{})
	logConfig["filename"] = gSecKillConf.LogPath
	logConfig["level"] = converLogLevel(gSecKillConf.LogLevel)
	logConfigStr, err := json.Marshal(logConfig)
	if err != nil {
		logs.Error("marshal failed, err:%v", err)
		return err
	}
	logs.SetLogger(logs.AdapterFile, string(logConfigStr))
	return
}

func initRedis() (err error) {

	gRedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", gSecKillConf.RedisConf.RedisAddr)
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         gSecKillConf.RedisConf.RedisMaxIdle,
		MaxActive:       gSecKillConf.RedisConf.RedisMaxActive,
		IdleTimeout:     time.Duration(gSecKillConf.RedisConf.RedisIdleTimeout) * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	conn := gRedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
	}
	return
}

func initEtcd() (err error) {
	gEtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{gSecKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(gSecKillConf.EtcdConf.EtcdTimeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:%v", err)
		return
	}
	return
}

func converLogLevel(level string) int {
	switch level {
	case "info":
		return logs.LevelInfo
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	}
	//LevelTrace = LevelDebug
	return logs.LevelTrace
}

// 加载秒杀商品信息
func loadSKProdInfos() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//resp, err := gEtcdClient.Get(ctx, "greeting")
	key := gSecKillConf.EtcdConf.EtcdProductKey
	resp, err := gEtcdClient.Get(ctx, key)
	if err != nil {
		logs.Error("get [%s] from etcd failed,err:%v", key, err)
		return
	}
	var products []service.SKProductInfo
	for _, v := range resp.Kvs {
		//logs.Debug("etcd resp is [%s]", v.Value)
		if err = json.Unmarshal(v.Value, &products); err != nil {
			logs.Error("json Unmarshal failed, err:%v", err)
			return
		}
		logs.Debug("seckill products info is = %v", products)
	}
	updateSKProdMap(products)
	go watchEtcd()
	return
}

// 监视etcd 是否发生变化
func watchEtcd() {
	key := gSecKillConf.EtcdConf.EtcdProductKey
	logs.Debug("begin watch key = %s", key)
	for {
		rch := gEtcdClient.Watch(context.Background(), key)
		var products []service.SKProductInfo
		getConfSucc := true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s]'s config is deleted", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					if err := json.Unmarshal(ev.Kv.Value, &products); err != nil {
						logs.Error("key[%s],Unmarshal failed,err:%v", key, err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, Type[%s],Key[%q],Value[%q]", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", products)
				updateSKProdMap(products)
			}
		}
	}

}

func updateSKProdMap(skProdVec []service.SKProductInfo) {
	tmpMap := make(map[int]*service.SKProductInfo, 1024)
	for _, v := range skProdVec {
		tmp := v
		tmpMap[v.ProductID] = &tmp
	}
	gSecKillConf.SKProdInfosRWLock.Lock()
	gSecKillConf.SKProdInfosMap = tmpMap
	gSecKillConf.SKProdInfosRWLock.Unlock()
}
