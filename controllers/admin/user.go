package admin

type UserController struct {
	baseController
}

// 用户列表
func (c *UserController) Index() {
	if c.IsAjax() {
		c.Data["json"] = map[string]interface{}{
			"code": 0, "msg": "", "count": "100",
			"data": map[string]interface{}{"cc": "222"},
		}
		c.ServeJSON()
		c.StopRun()
	}
	//var page int64
	//var pagesize int64 = 15
	//var list []*models.User
	//
	//if page, _ = c.GetInt64("page"); page < 1 {
	//	page = 1
	//}
	//
	//offset := (page - 1) * pagesize
	//
	//count, _ := user.Query().Count()
	c.display()
}

func (c *UserController) Create() {
	c.display()
}

func (c *UserController) Table() {
	c.Data["code"] = 0
	c.Data["msg"] = "请求成功"
	c.Data["count"] = 100
	c.ServeJSON()
}
