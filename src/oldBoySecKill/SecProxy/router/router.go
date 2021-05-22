package router

import (
	beego "github.com/astaxie/beego/adapter"
	"github.com/astaxie/beego/core/logs"
	"oldBoySecKill/SecProxy/controller"
)

func init() {
	logs.Debug("enter router init")
	beego.Router("/seckill", &controller.SecKillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SecKillController{}, "*:SecInfo")
}
