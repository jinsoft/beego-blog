package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id int64 `json:"id"`
	//User        *Users    `orm:"rel(fk)":index;json:"uid"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Views    int64     `json:"views"`
	Comments int64     `json:"comments"`
	Status   int8      `json:"status"`
	IsTop    int8      `json:"is_top"`
	Category *Category `orm:"rel(one)";json:"category"`
	//CategoryId  int       `json:"category_id"`
	CreatedTime time.Time `orm:"auto_now_add";json:"created_time"`
	UpdatedTime time.Time `orm:"auto_now_add";json:"updated_time"`
}

func (a *Article) TableName() string {
	return TableName("article")
}

func ArticleCreate(c *Article) (int64, error) {
	return orm.NewOrm().Insert(c)
}

func GetArticleList(page, pageSize int) ([]*Article, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(new(Article).TableName())
	total, _ := query.Count()
	query.Limit(pageSize, offset).All(&list)
	return list, total
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
