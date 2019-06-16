package index

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type List struct {
	limit int
	keyword string
}

func (c *IndexController) Index(){
	c.Ctx.WriteString(``)
	//id := c.GetString("id")
	//c.Ctx.WriteString(id + "\n")
	//c.Ctx.WriteString("index111111")
}