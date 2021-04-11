package main

import (
	_ "SecKill/Proxy/router"
	"github.com/astaxie/beego"
)

func main() {
	err := initConfig()
	if err !=nil{
		panic(err)
		return
	}
	err = initSec()
	if err !=nil{
		panic(err)
		return
	}
	beego.Run()
}
