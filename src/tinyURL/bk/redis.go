package bk

import (
	"context"
	"github.com/go-redis/redis"
)

type RedisCli struct {
	Cli *redis.Client
}

func NewReidsClient(addr string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
	if _, err := c.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}
