package admin

import "web/util"

type IndexController struct {
	baseController
}

func (c *IndexController) Index() {
	c.Data["username"] = c.username
	c.TplName = "admin/index.html"
}

func (c *IndexController) Dashboard() {
	c.TplName = "admin/dashboard.html"
}

func (c *IndexController) Test() {
	uid := util.GUID()
	c.Ctx.WriteString(uid)
}
