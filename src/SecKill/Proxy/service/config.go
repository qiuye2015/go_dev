package service

import (
	"sync"
)

type SecSkillConf struct{
	RedisConf         RedisConf
	EtcdConf          EtcdConf
	LogPath           string
	LogLevel          string
	SecProductInfoMap map[int]*SecProductInfo
	RWSecProductLock  sync.RWMutex
}
type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}
type EtcdConf struct{
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPreifx  string
	EtcdSecProductKey string
}

type SecProductInfo struct {
	ProductId int
	StartTime	int64
	EndTime	int64
	Status	int
	Total int
	Left int
}

const (
	ProductStatusNormal =0
	ProductStatusSaleOut = 1
	ProductStatusForceSaleOut = 2
)
