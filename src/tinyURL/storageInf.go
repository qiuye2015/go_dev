package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type StorageInf interface {
	Shorten(url string, exp int64) (string, error)
	UnShorten(eid string) (string, error)
	ShortLinkInfo(eid string) (string, error)
}

const (
	URLIdKey           = "next.url.id"         //全局自增器
	URLHashKey         = "urlhash:%s:url"      //地址hash和短地址映射
	ShortLinkKey       = "shortlink:%s:url"    //短地址和原地址的映射
	ShortLinkDetailKey = "shortlink:%s:detail" //短地址和详情的映射
)

type RedisCli struct {
	Cli *redis.Client
}

//短地址详细信息
type URLDetail struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"create_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

func NewRedisCli(addr string, db int) RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
		//Network:            "",
		//Dialer:             nil,
		//OnConnect:          nil,
		//Username:           "",
		//Password:           "",
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		//DialTimeout:        0,
		//ReadTimeout:        0,
		//WriteTimeout:       0,
		//PoolSize:           0,
		//MinIdleConns:       0,
		//MaxConnAge:         0,
		//PoolTimeout:        0,
		//IdleTimeout:        0,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
		//Limiter:            nil,
	})
	if _, err := c.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return RedisCli{Cli: c}
}

func (r RedisCli) Shorten(url string, exp int64) (string, error) {
	//1. 计算长地址的hash,便于存储长地址作key
	urlHash := toSha1(url)
	ctx := context.Background()
	res, err := r.Cli.Get(ctx, fmt.Sprintf(URLHashKey, urlHash)).Result()
	var id int64
	if err == redis.Nil { //URLHashKey不存在,取自增量ID
		id, err = r.Cli.Incr(ctx, URLIdKey).Result()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else { //存在这个url对应的短地址
		if res != "{}" {
			id = Base62Decode(res)
			return res, nil
		}
	}
	//2. 把id转成62进制
	encodeId := Base62Encode(id)
	log.Printf("encodeId: %v", encodeId)
	//3. 存短地址和长地址的映射
	err = r.Cli.Set(ctx, fmt.Sprintf(ShortLinkKey, encodeId), url, time.Duration(exp)*time.Minute).Err()
	if err != nil {
		return "", err
	}
	//4. 存长地址哈希值和短地址的映射
	err = r.Cli.Set(ctx, fmt.Sprintf(URLHashKey, urlHash), encodeId, time.Duration(exp)*time.Minute).Err()
	if err != nil {
		return "", err
	}
	detail, err := json.Marshal(&URLDetail{
		URL:                 url,
		CreateAt:            time.Now().String(),
		ExpirationInMinutes: time.Duration(exp),
	})
	if err != nil {
		return "", err
	}
	//5. 存短地址和详情的映射
	err = r.Cli.Set(ctx, fmt.Sprintf(ShortLinkDetailKey, encodeId), detail, time.Duration(exp)*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return encodeId, nil
}

//获取长地址
func (r RedisCli) UnShorten(encodeId string) (string, error) {
	ctx := context.Background()
	url, err := r.Cli.Get(ctx, fmt.Sprintf(ShortLinkKey, encodeId)).Result()
	if err == redis.Nil {
		return "", StatusError{
			Code: 404,
			Err:  errors.New("unknow short URL"),
		}
	} else if err != nil {
		return "", err
	}
	return url, nil
}

// 获取短地址的详细信息
func (r RedisCli) ShortLinkInfo(eid string) (string, error) {
	ctx := context.Background()
	detail, err := r.Cli.Get(ctx, fmt.Sprintf(ShortLinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{
			Code: 404,
			Err:  errors.New("unknow short URL"),
		}
	} else if err != nil {
		return "", err
	}
	return detail, nil
}
