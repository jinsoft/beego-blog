package admin

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
