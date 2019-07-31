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
	beego.Router("/admin/test", &admin.IndexController{}, "get:Test")
	beego.Router("/admin/login", &admin.AuthController{}, "get:ShowLoginForm;post:Login")
	beego.Router("/admin", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/dashboard", &admin.IndexController{}, "get:Dashboard")
	// admin user module
	beego.Router("/admin/user/index", &admin.UserController{}, "get:Index")
	beego.Router("/admin/user/create", &admin.UserController{}, "get:Create;post:Create")
	beego.Router("/admin/user/:id:int", &admin.UserController{}, "get:Edit;post:Edit;delete:Destroy")
	// admin article module
	beego.Router("/admin/article/index", &admin.ArticleController{}, "get:Index")
	beego.Router("/admin/category/index", &admin.CategoryController{}, "get:Index;post:Index")
	beego.Router("/admin/category/create", &admin.CategoryController{}, "get:Create;post:Create")
	beego.Router("/admin/category/edit/:id:int", &admin.CategoryController{}, "get:Edit;post:Edit")
	beego.Router("/admin/category/:id:int", &admin.CategoryController{}, "post:Destroy")
	// admin common func
	beego.Router("/admin/upload", &admin.CommonController{}, "post:Upload")
}
