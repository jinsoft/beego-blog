package main

import (
	"github.com/astaxie/beego"
	_ "web/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
