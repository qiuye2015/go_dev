package service

import (
	"github.com/gomodule/redigo/redis"
	clientV3 "go.Etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type SecKillCtx struct {
	LogPath  string
	LogLevel string

	RedisConf RedisConfig
	EtcdConf  EtcdConfig

	RedisPool  *redis.Pool
	EtcdClient *clientV3.Client

	SKProdInfosMap    map[int]*SKProductInfo
	SKProdInfosRWLock sync.RWMutex

	CookieSecretKey            string
	UserAccessLimitEverySecond int
	IpAccessLimitEverySecond   int
	ReferWhiteList             []string

	IdBlackMap       map[int]bool
	IpBlackMap       map[string]bool
	IdBlackMapRWLock sync.RWMutex
	IpBlackMapRWLock sync.RWMutex

	SecKillReqChan    chan *SKRequest
	SecKillReqChanNum int
}
type RedisConfig struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}
type EtcdConfig struct {
	EtcdAddr       string
	EtcdTimeout    int
	EtcdKeyPrefix  string
	EtcdProductKey string
}

//SecKillProductInfo 秒杀商品信息
type SKProductInfo struct {
	ProductID int
	Total     int // 总量
	Stock     int // 库存
	StartTime int64
	EndTime   int64
	Status    int
}

const (
	ProductStatusNormal = iota
	ProductStatusSaleOut
	ProductStatusForceSaleOut
)

type SKRequest struct {
	ProductID    int
	Source       string
	AuthCode     string
	SecTime      string
	Nance        string
	UserID       int
	UserAuthSign string
	AccessTime   time.Time

	ClientAddr  string
	ClientRefer string

	CloseNotify <-chan bool
	ResultChan  chan *SKResult
}
type SKResult struct {
}
