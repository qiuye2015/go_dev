package main

import (
	"fmt"
	beego "github.com/astaxie/beego"
	"log"
	//beego "github.com/astaxie/beego/adapter"
)

type RESTfulController struct {
	beego.Controller
}

func (this *RESTfulController) Get() {
	this.Ctx.WriteString("Hello World in GET method!")
}

func (this *RESTfulController) Post() {
	this.Ctx.WriteString("Hello World in POST method!")
}

//beego支持正则表达式,gin不持支
func main_beego(host string) {
	// RESTful Controller 路由
	beego.Router("/RESTful", &RESTfulController{})
	//正则路由
	beego.Router("/RegExp1/?:id", &RegExpfulController{})
	beego.Router("/RegExp2/:id([0-9]+)", &RegExpfulController{})
	beego.Router("/RegExp3/:id([\\w]+)", &RegExpfulController{})
	beego.Router("/RegExp4/abc:id([0-9]+)de", &RegExpfulController{})
	beego.Router("/RegExp5/*", &RegExpfulController{})
	beego.Router("/RegExp6/*.*", &RegExpfulController{})

	beego.Run(host)
}

type RegExpfulController struct {
	beego.Controller
}

func (this *RegExpfulController) Get() {
	this.Ctx.WriteString(fmt.Sprintf("In RegExp Mode!\n"))
	id := this.Ctx.Input.Param(":id")
	log.Println(this.Ctx.Request.URL)
	this.Ctx.WriteString(fmt.Sprintf("id is :%s\n", id))

	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString(fmt.Sprintf("splat is :%s\n", splat))

	path := this.Ctx.Input.Param(":path")
	this.Ctx.WriteString(fmt.Sprintf("path is :%s\n", path))

	ext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString(fmt.Sprintf("ext is :%s\n", ext))
}
