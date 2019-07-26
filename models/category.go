package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id           int64
	CategoryName string
	Order        int8
	CreateTime   time.Time
}

func (c *Category) TableName() string {
	return TableName("category")
}

func GetCategoryList(page, pageSize int) ([]*Category, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Category, 0)
	query := orm.NewOrm().QueryTable(new(Category).TableName())
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}
