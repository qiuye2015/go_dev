package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	secKillConf *SecSkillConf
)
func InitService(serviceConf *SecSkillConf){
	secKillConf = serviceConf
	logs.Debug("init service succ,config:%v",secKillConf)
}

func SecInfo(productID int)(data map[string]interface{},code int,err error){
	//data=make(map[string]interface{},16)
	//data["sec"] = "test"
	//data["begin"] = time.Now()
	secKillConf.RWSecProductLock.Lock()
	defer secKillConf.RWSecProductLock.Unlock()

	v,ok := secKillConf.SecProductInfoMap[productID]
	if !ok {
		code = ErrNotFoundProductID
		err = fmt.Errorf("not found product_id %v",productID)
		return
	}
	start:=false
	end:=false
	status:="success"
	now:=time.Now().Unix()
	if now -v.StartTime >0 {
		start =true
	}
	if now-v.EndTime>0{
		start=false
		end =true
	}
	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut{
		start = false
		end = true
		status = "prodcut is sale out"
	}

	data=make(map[string]interface{},16)
	data["product_id"]=productID
	data["start_time"] = start
	data["end_time"] = end
	data["status"] =status

	return
}

