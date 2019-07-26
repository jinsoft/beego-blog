package admin

import (
	"html/template"
	"web/models"
)

type CategoryController struct {
	baseController
}

func (c *CategoryController) Index() {
	if c.IsAjax() {
		page, err := c.GetInt("page")
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := c.GetInt("limit")
		if err != nil {
			limit = 15
		}
		result, count := models.GetCategoryList(page, limit)
		c.Data["json"] = map[string]interface{}{
			"code":  0,
			"msg":   "请求成功",
			"count": count,
			"data":  result,
		}
		c.ServeJSON()
		c.StopRun()
	}
	c.display()
}

func (c *CategoryController) Create() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.IsAjax() {
		//
	}
	c.display()
}
