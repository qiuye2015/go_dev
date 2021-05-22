package service

import (
	"sync"
)

type SecKillConf struct {
	LogPath  string
	LogLevel string

	RedisConf RedisConfig
	EtcdConf  EtcdConfig

	SKProdInfosMap    map[int]*SKProductInfo
	SKProdInfosRWLock sync.RWMutex

	CookieSecretKey string
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
}
