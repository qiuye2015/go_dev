package main

import (
	"context"
	"encoding/json"
	"fmt"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"github.com/qiuye2015/go_dev/oldBoySecKill/SecProxy/service"
	//"go.etcd.io/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/api/v3/mvccpb"

	clientV3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

var (
	gSecKillConf = &service.SecKillCtx{
		SKProdInfosMap: make(map[int]*service.SKProductInfo, 1024),
	}
)

func initConfig() (err error) {
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

	gSecKillConf.CookieSecretKey = beego.AppConfig.String("dev::cookieSecretKey")
	userAccessLimitEverySecond, err := beego.AppConfig.Int("dev::userAccessLimitEverySecond")
	if err != nil {
		err = fmt.Errorf("init config Int failed,err:%v", err)
		return
	}
	gSecKillConf.UserAccessLimitEverySecond = userAccessLimitEverySecond

	ipAccessLimitEverySecond, err := beego.AppConfig.Int("dev::ipAccessLimitEverySecond")
	if err != nil {
		err = fmt.Errorf("init config Int failed,err:%v", err)
		return
	}
	gSecKillConf.IpAccessLimitEverySecond = ipAccessLimitEverySecond

	refers := beego.AppConfig.String("dev::referWhiteList")
	if len(refers) > 0 {
		gSecKillConf.ReferWhiteList = strings.Split(refers, ",")
	}
	logs.Info("initConfig success gSecKillConf = %v", gSecKillConf)
	return
}

func initSecKill() (err error) {
	if err = initLogs(); err != nil {
		logs.Error("init logger failed, err:%v", err)
		return
	}
	if err = initRedis(); err != nil {
		logs.Error("init redis failed, err:%v", err)
		return
	}
	if err = initEtcd(); err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	if err = loadSKProdInfos(); err != nil {
		logs.Error("loadSKProdInfos failed, err:%v", err)
		return
	}

	logs.Info("init sec kill success")
	return
}

func initLogs() (err error) {
	logConfig := make(map[string]interface{})
	logConfig["filename"] = gSecKillConf.LogPath
	logConfig["level"] = convertLogLevel(gSecKillConf.LogLevel)
	logConfigStr, err := json.Marshal(logConfig)
	if err != nil {
		logs.Error("marshal failed, err:%v", err)
		return err
	}
	_ = logs.SetLogger(logs.AdapterFile, string(logConfigStr))
	return
}

func initRedis() (err error) {

	gRedisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", gSecKillConf.RedisConf.RedisAddr)
		},
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
		return
	}
	gSecKillConf.RedisPool = gRedisPool
	return
}

func initEtcd() (err error) {
	gEtcdClient, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{gSecKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(gSecKillConf.EtcdConf.EtcdTimeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:%v", err)
		return
	}
	gSecKillConf.EtcdClient = gEtcdClient
	return
}

func convertLogLevel(level string) int {
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
	resp, err := gSecKillConf.EtcdClient.Get(ctx, key)
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
		logs.Debug("sec kill products info is = %v", products)
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
		rch := gSecKillConf.EtcdClient.Watch(context.Background(), key)
		var products []service.SKProductInfo
		getConfSuccess := true
		for resp := range rch {
			for _, ev := range resp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s]'s config is deleted", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					if err := json.Unmarshal(ev.Kv.Value, &products); err != nil {
						logs.Error("key[%s],Unmarshal failed,err:%v", key, err)
						getConfSuccess = false
						continue
					}
				}
				logs.Debug("get config from etcd, Type[%s],Key[%q],Value[%q]", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSuccess {
				logs.Debug("get config from etcd success, %v", products)
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
