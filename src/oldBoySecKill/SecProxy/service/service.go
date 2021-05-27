package service

import (
	"fmt"
	"github.com/astaxie/beego/core/logs"
	"time"
)

var (
	gServiceConf *SecKillCtx
)

func SecInfo(productID int) (data []map[string]interface{}, code int, err error) {
	gServiceConf.SKProdInfosRWLock.RLock()
	//defer gServiceConf.SKProdInfosRWLock.RUnlock()
	v, ok := gServiceConf.SKProdInfosMap[productID]
	gServiceConf.SKProdInfosRWLock.RUnlock()
	if !ok {
		code = ErrCodeNotFoundProdutID
		err = fmt.Errorf("not found product_id:%d", productID)
		return
	}
	start := false
	status := ""
	now := time.Now().Unix()

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut {
		status = "product is sale out"
		code = ErrCodeActiveSaleOut
	} else if (v.Status == ProductStatusNormal) && now >= v.StartTime && now <= v.EndTime {
		start = true
		status = "sec kill is starting"
		code = ErrCodeNormal
	} else if now < v.EndTime {
		status = "sec kill is not start"
		code = ErrCodeActiveNotStart
	} else {
		status = "sec kill is already end"
		code = ErrCodeActiveAlreadyEnd
	}

	item := make(map[string]interface{})
	item["product_id"] = productID
	item["start_flag"] = start
	item["status"] = status
	data = append(data, item)
	return
}

//显示所有秒杀商品信息
func SecInfoList() (data []map[string]interface{}, code int, err error) {
	gServiceConf.SKProdInfosRWLock.RLock()
	defer gServiceConf.SKProdInfosRWLock.RUnlock()
	//result := make([]map[string]interface{}, 0, 10)

	for _, v := range gServiceConf.SKProdInfosMap {
		items, _, err := SecInfo(v.ProductID)
		if err != nil {
			logs.Warn("get product info failed,err:%v", err)
			continue
		}
		data = append(data, items...)
	}
	return
}

func SecKill(req *SKRequest) (data []map[string]interface{}, code int, err error) {
	if err = userCheck(req); err != nil {
		code = ErrCodeUserCheckFailed
		logs.Warn("userID = %d check failed, req = %v", req.UserID, req)
		return
	}
	if err = antiSpam(req); err != nil {
		code = ErrCodeServiceBusy
		logs.Warn("userID = %d check failed, req = %v", req.UserID, req)
		return
	}
	data, code, err = SecInfo(req.ProductID)
	if err != nil || code != ErrCodeNormal {
		logs.Warn("userID[%d],get Sec Info failed, code:%v, err:%v", req.UserID, code, err)
		return
	}
	gServiceConf.SecKillReqChan <- req
	return
}
