package admin

import (
	"strconv"
	"strings"
	"web/models"
)

type ArticleController struct {
	baseController
}

type article struct {
	Id       int64
	Title    string `valid:"Required"`
	Status   int8   `valid:"Required"`
	Category int64  `valid:"Required"`
}

//
//func (a *article) Valid(v *validation.Validation){
//}

func (c *ArticleController) Index() {
	if c.IsAjax() {
		result, count := models.GetArticleList(c.pageNumber, c.pageSize)
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

func (c *ArticleController) Create() {
	if c.IsAjax() {
		//name := strings.TrimSpace(c.GetString("md-content-html-code"))
		title := strings.TrimSpace(c.GetString("title"))
		content := strings.TrimSpace(c.GetString("content-html-code"))
		category := strings.TrimSpace(c.GetString("category"))
		status, err := c.GetInt8("status")
		if err != nil {
			status = 0
		}
		category_id, _ := strconv.Atoi(category)

		Article := new(models.Article)
		Article.Uid = c.uid
		Article.Content = content
		Article.Title = title
		Article.Status = status
		Article.Cid = category_id
		if _, err := models.ArticleCreate(Article); err != nil {
			c.ajaxMsg(err.Error(), MSG_ERR)
		}
		c.ajaxMsg("添加成功", MSG_OK)
	}
	// 获取所有的分类
	category := models.GetCategoryAllList()
	c.Data["category"] = category
	c.display()
}

func (c *ArticleController) Edit() {
	c.display()
}

func (c *ArticleController) Destroy() {
	//
}
