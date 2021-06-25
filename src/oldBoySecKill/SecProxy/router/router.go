package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/qiuye2015/go_dev/oldBoySecKill/SecProxy/controller"
	//"github.com/qiuye2015/go_dev/oldBoySecKill/SecProxy/controller"
)

func init() {
	logs.Debug("enter router init")
	beego.Router("/seckill", &controller.SecKillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SecKillController{}, "*:SecInfo")
}
