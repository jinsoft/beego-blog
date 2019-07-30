package admin

import (
	"strconv"
	"strings"
	"web/models"
)

type User struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}

type AuthController struct {
	baseController
}

func (c *AuthController) Prepare() {
	c.EnableXSRF = true
}

// 登录页
func (c *AuthController) ShowLoginForm() {
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "admin/login.html"
}

// 处理登录请求
func (c *AuthController) Login() {
	email := strings.TrimSpace(c.GetString("email"))
	password := strings.TrimSpace(c.GetString("password"))
	if email != "" && password != "" {
		var admin models.Admins
		admin.Email = email
		//admin.Email = email
		if admin.Read("email") != nil || admin.Password != models.Md5([]byte(password)) {
			c.Data["json"] = map[string]interface{}{"code": 401, "message": "帐号或密码错误"}
		} else {
			authkey := models.Md5([]byte(c.getClientIp() + "|" + admin.Password))
			c.Ctx.SetCookie("ainiok_session", strconv.FormatInt(admin.Id, 10)+"|"+authkey)
			//c.Redirect("/admin", 302)
			c.Data["json"] = map[string]interface{}{"code": 200, "message": "登录成功", "url": "/admin"}
		}
	} else if email == "" {
		c.Data["json"] = map[string]interface{}{"code": 401, "message": "邮箱不能为空"}
	} else if password == "" {
		c.Data["json"] = map[string]interface{}{"code": 401, "message": "密码不能为空"}
	}
	c.ServeJSON()
}
