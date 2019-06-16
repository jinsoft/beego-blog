package admin

import (
	"github.com/astaxie/beego"
	"net"
	"strings"
)

type baseController struct {
	beego.Controller
	uid            int64
	controllerName string
	actionName     string
	clientIp       string
}

func (c *baseController) Prepare() {
	c.clientIp = c.getClientIp()
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	c.auth()
}

func (c *baseController) auth() {
	//if c.uid == 0 && c.actionName != "showloginform" && c.actionName != "login" {
	//	c.Redirect("/", 302)
	//}
}

// 获取用户ip地址
func (c *baseController) getClientIp() string {
	ip := strings.TrimSpace(strings.Split(c.Ctx.Request.Header.Get("X-Forwarded-For"),",")[0])
	if ip != "" {
		return ip
	}
	ip = c.Ctx.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Ctx.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
