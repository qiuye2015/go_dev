package main

import (
	beego "github.com/astaxie/beego"
	_ "github.com/qiuye2015/go_dev/oldBoySecKill/SecProxy/router"
	"github.com/qiuye2015/go_dev/oldBoySecKill/SecProxy/service"
)

func main() {
	//str, _ := os.Getwd()
	//fmt.Println(str)
	if err := initConfig(); err != nil {
		panic(err)
		return
	}
	if err := initSecKill(); err != nil {
		panic(err)
		return
	}
	service.InitService(gSecKillConf)
	beego.Run()
}
