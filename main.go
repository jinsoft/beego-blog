package main

import (
	"github.com/astaxie/beego"
	"web/controllers/admin"
	_ "web/routers"
)

func main() {
	beego.ErrorController(&admin.ErrorController{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
