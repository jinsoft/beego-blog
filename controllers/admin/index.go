package admin

type IndexController struct {
	baseController
}

func (c *IndexController) Index(){
	c.TplName = "admin/index.html"
}