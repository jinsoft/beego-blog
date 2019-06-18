package admin

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}

type AuthController struct {
	baseController
}

type user struct {
	Username string `valid:"Required;Email;MaxSize(100)"`
	Password string `valid:MaxSize(100)`
}

func (u *user) Valid(v *validation.Validation) {
	//
}

func (c *AuthController) ShowLoginForm() {
	c.Data["xsrfdata"] = c.XSRFToken()
	c.TplName = "admin/login.tpl"
}

func (c *AuthController) Login() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("error")
	}
	if u.Name == "admin" && u.Password == "123456" {
		c.uid = 1
		c.Data["json"] = map[string]interface{}{"code": 200, "message": "登录成功", "url": "/admin"}
		//c.Redirect("/admin", 302)
	} else {
		c.Data["json"] = map[string]interface{}{"code": 401, "message": "登录失败"}
	}
	c.ServeJSON()
}
