package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Addr string
	Db   int
}

//在环境变量中初始化配置文件
func (c *Config) InitConfig() {
	c.Addr = os.Getenv("APP_REDIS_ADDR")
	if len(c.Addr) == 0 {
		c.Addr = "localhost:6379"
	}
	dbs := os.Getenv("APP_REDIS_DB")
	if len(dbs) == 0 {
		dbs = "0"
	}
	db, err := strconv.Atoi(dbs)
	if err != nil {
		log.Fatal(err)
	}
	c.Db = db
	log.Printf("connect to redis (Addr: %s Db: %d)", c.Addr, c.Db)
}
