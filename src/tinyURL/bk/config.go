package bk

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	RedisAddr string
	RedisDb   int
	Cli       *RedisCli
}

func InitConfig() *Config {
	addr := os.Getenv("APP_REDIS_ADDR")
	if len(addr) == 0 {
		addr = "localhost:6379"
	}

	dbStr := os.Getenv("APP_REDIS_DB")
	if len(dbStr) == 0 {
		dbStr = "0"
	}
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0
	}
	log.Printf("connect to redis (addr: %s, db: %d)", addr, db)

	return &Config{
		RedisAddr: addr,
		RedisDb:   db,
		Cli:       NewReidsClient(addr, db),
	}
}
