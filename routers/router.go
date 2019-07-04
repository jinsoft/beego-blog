package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
	"web/controllers/admin"
	"web/controllers/index"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/index", &index.IndexController{}, "get:Index")

	// admin routers
	beego.Router("/admin/login", &admin.AuthController{}, "get:ShowLoginForm;post:Login")
	beego.Router("/admin", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/dashboard", &admin.IndexController{}, "get:Dashboard")
	beego.Router("/admin/user/index", &admin.UserController{}, "get:Index")
	beego.Router("/admin/user/create", &admin.UserController{}, "get:Create")

	// admin common func
	beego.Router("/admin/upload", &admin.CommonController{}, "post:Upload")
}
