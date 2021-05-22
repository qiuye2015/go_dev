package main

import (
	beego "github.com/astaxie/beego/adapter"
	_ "oldBoySecKill/SecProxy/router"
	"oldBoySecKill/SecProxy/service"
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
