package controller

import (
	beego "github.com/astaxie/beego/adapter"
	"github.com/astaxie/beego/core/logs"
	"oldBoySecKill/SecProxy/service"
)

type SecKillController struct {
	beego.Controller
}

func (p *SecKillController) SecKill() {
	result := make(map[string]interface{})
	result["code"] = service.ErrCodeNormal
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	//var data []map[string]interface{}
	//var code int
	productID, err := p.GetInt("product_id")
	if err != nil {
		result["code"] = service.ErrCodeInvalidRequest
		result["message"] = err.Error()
		logs.Error("invaild request,get product_id failed,err:%v", err)
		return
	}
	secRequest := &service.SKRequest{}
	secRequest.ProductID = productID
	secRequest.Source = p.GetString("source")     // 来源
	secRequest.AuthCode = p.GetString("authcode") // 鉴权码
	secRequest.SecTime = p.GetString("time")      // 当前时间
	secRequest.Nance = p.GetString("nance")       // 随机数
	secRequest.UserID = p.Ctx.GetCookie("userID")
	secRequest.UserAuthSign = p.Ctx.GetCookie("userAuthSign")

	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["message"] = err.Error()
		logs.Error("invaild request,get product_id failed,err:%v", err)
		return
	}
	result["code"] = code
	result["data"] = data
}

func (p *SecKillController) SecInfo() {
	result := make(map[string]interface{})
	result["code"] = service.ErrCodeNormal
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	var data []map[string]interface{}
	var code int
	productID, err := p.GetInt("product_id")

	if err == nil {
		data, code, err = service.SecInfo(productID)
	} else {
		data, code, err = service.SecInfoList()
	}

	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		logs.Error("invaild request,get product_id failed,err:%v", err)
	} else {
		result["data"] = data
	}

}
