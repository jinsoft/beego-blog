package admin

import (
	"github.com/astaxie/beego/orm"
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
		//var (
		//	title     string
		//	aid int64
		//	tag       string
		//)
		//
		//title = c.GetString("title")
		//aid,_ = c.GetInt64("aid")
		//tag = c.GetString("tag")
		var (
			list    []*models.Article
			article models.Article
		)

		articleid, _ := c.GetInt64("article_id")

		if articleid != 0 {
			article.Query().Filter("id", articleid).One(&article)
		}

		query := article.Query()
		count, _ := query.Count()
		if count > 0 {
			query.OrderBy("-istop", "-id").Limit(c.pageSize, c.offset).RelatedSel().All(&list)
		}

		data := make(map[int]interface{})
		for k, v := range list {
			row := make(map[string]interface{})
			row["id"] = v.Id
			row["title"] = v.Title
			row["content"] = v.Content
			row["comments"] = v.Comments
			row["views"] = v.Views
			row["is_top"] = v.IsTop
			row["status"] = v.Status
			row["is_top"] = v.IsTop
			row["category_name"] = v.Category.CategoryName
			if v.UpdatedTime.IsZero() {
				row["updated_time"] = "-"
			} else {
				row["updated_time"] = v.UpdatedTime.Format("2006-01-02 15:04")
			}
			row["created_time"] = v.CreatedTime.Format("2006-01-02 15:04")
			data[k] = row
		}

		//result, count := models.GetArticleList(c.pageNumber, c.pageSize)

		c.Data["json"] = map[string]interface{}{
			"code":  0,
			"msg":   "请求成功",
			"count": count,
			"data":  data,
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
		tags := strings.TrimSpace(c.GetString("tags"))
		status, err := c.GetInt8("status")
		is_top, err := c.GetInt8("is_top")
		if err != nil {
			status = 0
		}
		category_id, _ := strconv.ParseInt(category, 0, 64)

		addTags := make([]string, 0)
		// 标签过滤
		if tags != "" {
			tagArr := strings.Split(tags, ",")
			for _, v := range tagArr {
				if tag := strings.TrimSpace(v); tag != "" {
					exists := false
					for _, vv := range addTags {
						if vv == tag {
							exists = true
							break
						}
					}
					if !exists {
						addTags = append(addTags, tag)
					}
				}
			}
		}

		if len(addTags) > 0 {
			for _, v := range addTags {
				tag := models.Tags{Name: v}
				if tag.Read("Name") == orm.ErrNoRows {
					tag.Count = 1
					tag.Insert()
				} else {
					tag.Count += 1
					tag.Update("Count")
				}
			}
		}

		Article := new(models.Article)
		//Article.User = &models.Users{Id: c.userid}
		Article.Content = content
		Article.Title = title
		Article.Status = status
		Article.IsTop = is_top
		//Article.CategoryId = category_id
		Article.Category = &models.Category{Id: category_id}
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
