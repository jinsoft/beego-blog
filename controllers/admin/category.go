package admin

import (
	"html/template"
	"strconv"
	"strings"
	"time"
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
		name := strings.TrimSpace(c.GetString("name"))
		order, err := c.GetInt8("order")
		if err != nil {
			order = 99
		}
		Category := new(models.Category)
		Category.CategoryName = name
		Category.Order = order
		Category.CreateTime = time.Now()
		if _, err := models.CategoryAdd(Category); err != nil {
			c.ajaxMsg(err.Error(), MSG_ERR)
		}
		c.ajaxMsg("添加成功", MSG_OK)
	}
	c.display()
}

func (c *CategoryController) Edit() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	id := c.Ctx.Input.Param(":id")
	category_id, _ := strconv.ParseInt(id, 10, 64)
	if c.IsAjax() {
		Category, _ := models.GetCategoryById(category_id)
		Category.Id = category_id
		Category.CategoryName = strings.TrimSpace(c.GetString("name"))
		Category.Order, _ = c.GetInt8("order", 99)
		Category.UpdateTime = time.Now()
		if err := Category.Update(); err != nil {
			c.ajaxMsg(err.Error(), MSG_ERR)
		}
		c.ajaxMsg("修改成功", MSG_OK)
	}
	detail, _ := models.GetCategoryById(category_id)
	row := make(map[string]interface{})
	row["id"] = detail.Id
	row["name"] = detail.CategoryName
	row["order"] = detail.Order
	c.Data["Info"] = row
	c.display()
}

func (c *CategoryController) Destroy() {
	id := c.Ctx.Input.Param(":id")
	//category_id, _:=strconv.ParseInt(id,10,64)
	if c.IsAjax() {
		c.Ctx.WriteString(id)
	}
}
