package admin

type ArticleController struct {
	baseController
}

func (c *ArticleController) Index() {
	c.display()
}
