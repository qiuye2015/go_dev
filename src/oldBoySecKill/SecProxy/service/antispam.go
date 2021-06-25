package service

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
	"time"
)

var secLimitMgr = &SecLimitMgr{
	UserLimitMap: make(map[int]*SecLimit, 10000),
	IPLimitMap:   make(map[string]*SecLimit, 10000),
}

type SecLimitMgr struct {
	UserLimitMap map[int]*SecLimit
	IPLimitMap   map[string]*SecLimit
	lock         sync.Mutex
}
type SecLimit struct {
	count    int
	lastTime int64
}

func (p *SecLimit) Count(nowTime int64) int {
	if p.lastTime != nowTime {
		p.count = 1
		p.lastTime = nowTime
	} else {
		p.count++
	}
	return p.count
}
func (p *SecLimit) Check(nowTime int64) int {
	if p.lastTime != nowTime {
		return 0
	} else {
		return p.count
	}
}

//反作弊
func antiSpam(req *SKRequest) (err error) {
	secLimitMgr.lock.Lock()
	defer secLimitMgr.lock.Unlock()
	secLimit, ok := secLimitMgr.UserLimitMap[req.UserID]
	if !ok {
		secLimitMgr.UserLimitMap[req.UserID] = &SecLimit{}
	}
	count := secLimit.Count(time.Now().Unix())
	if count > gServiceConf.UserAccessLimitEverySecond {
		err = fmt.Errorf("invaild request")
		return
	}

	ipLimit, ok := secLimitMgr.IPLimitMap[req.ClientAddr]
	if !ok {
		secLimitMgr.IPLimitMap[req.ClientAddr] = &SecLimit{}
	}
	count = ipLimit.Count(time.Now().Unix())
	if count > gServiceConf.UserAccessLimitEverySecond {
		err = fmt.Errorf("invaild request")
		return
	}

	return
}

func userCheck(req *SKRequest) (err error) {
	found := false
	for _, v := range gServiceConf.ReferWhiteList {
		if v == req.ClientRefer {
			found = true
			break
		}
	}
	if !found {
		err = fmt.Errorf("invalid request")
		logs.Warn("user = %d is reject by refer,req = %v", req.UserID, req)
	}
	authData := fmt.Sprintf("%d%s", req.UserID, gServiceConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid user cookie auth")
	}
	return
}
