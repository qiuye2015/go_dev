package service

import (
	"encoding/json"
	"github.com/astaxie/beego/core/logs"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func InitService(serviceConf *SecKillCtx) {
	gServiceConf = serviceConf
	//logs.Debug("init service success, config = %v", serviceConf)
	if err := loadBlackList(); err != nil {
		logs.Error("load black List failed, err:%v", err)
		return
	}
	gServiceConf.SecKillReqChan = make(chan *SKRequest, gServiceConf.SecKillReqChanNum)
	initRedisProcess()
	return
}

func loadBlackList() (err error) {
	conn := gServiceConf.RedisPool.Get()
	defer conn.Close()
	reply, err := conn.Do("hgetall", "idBlackList")
	idLists, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hgetall failed,err:%v", err)
		return
	}
	for _, v := range idLists {
		id, err := strconv.Atoi(v)
		if err != nil {
			logs.Warn("invalid user id = %v", v)
			continue
		}
		gServiceConf.IdBlackMapRWLock.Lock()
		gServiceConf.IdBlackMap[id] = true
		gServiceConf.IdBlackMapRWLock.Unlock()
	}

	reply, err = conn.Do("hgetall", "ipBlackList")
	ipLists, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hgetall failed,err:%v", err)
		return
	}
	for _, v := range ipLists {
		gServiceConf.IpBlackMap[v] = true
	}
	go func() {
		for {
			reply, err := conn.Do("LPOP", "idBlackList")
			id, err := redis.Int(reply, err)
			if err != nil {
				continue
			}
			gServiceConf.IdBlackMap[id] = true
		}
	}()
	go func() {
		var ipList []string
		lastTime := time.Now().Unix()
		for {
			reply, err := conn.Do("LPOP", "ipBlackList")
			ip, err := redis.String(reply, err)
			if err != nil {
				continue
			}
			ipList = append(ipList, ip)
			curTime := time.Now().Unix()
			if len(ipList) > 100 || curTime-lastTime > 10 {
				gServiceConf.IpBlackMapRWLock.Lock()
				for _, v := range ipList {
					gServiceConf.IpBlackMap[v] = true
				}
				gServiceConf.IpBlackMapRWLock.Unlock()
				logs.Info("current time = %v, sync ip list success,ip [%v]", curTime, ipList)
				lastTime = curTime
				ipList = ipList[0:0]
			}

		}
	}()
	return
}

func initRedisProcess() {
	//for i := 0; i < 16; i++ {
	//	go writeHandler()
	//}
	//for i := 0; i < 16; i++ {
	//	go readHandler()
	//}
	return
}
func writeHandler() {

	for {
		conn := gServiceConf.RedisPool.Get()
		req := <-gServiceConf.SecKillReqChan
		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json Marshal failed,err:%v", err)
			continue
		}

		_, err = conn.Do("LPUSH", "secKillQueue", data)
		if err != nil {
			logs.Error("lpush failed,err%v,req:%v", err, req)
			continue
		}
		conn.Close()
	}
}

//TODO:
func readHandler() {
	for {
		conn := gServiceConf.RedisPool.Get()
		for {
			data, err := redis.String(conn.Do("RPOP", "secKillQueue"))
			if err != nil {
				logs.Error("rpop failed,err:%v, data:%v", err, data)
				break
			}
			logs.Debug("lpop from queue,data:%v", data)
			var seq SKRequest
			err = json.Unmarshal([]byte(data), &seq)
			if err != nil {
				logs.Error("json unmarshal failed,err:%v", err)
				continue
			}
			now := time.Now().Unix()
			if now-seq.AccessTime.Unix() > 10 {
				logs.Warn("req[%v] is expire", seq)
				continue
			}
			//<-seq
		}
		conn.Close()
	}
}
