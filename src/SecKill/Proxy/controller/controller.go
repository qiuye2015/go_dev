package controller

import (
	"SecKill/Proxy/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	p.Data["json"] = "sec kill"
	p.ServeJSON()
}

func (p *SkillController) SecInfo() {
	result:=make(map[string]interface{})
	result["code"] = 0
	result["message"] = "succ"
	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	prd,err:=p.GetInt("product_id")
	if err !=nil{
		result["code"] = -1
		result["message"] = "invaild product_id"
		logs.Error("invaild request,get product_id failed,err:%v",err)
		return
	}
	data,code,err:= service.SecInfo(prd)
	if err !=nil{
		result["code"] = code
		result["message"] = err.Error()
	}
	result["data"] = data
}
