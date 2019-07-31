package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	Order        int8   `json:"order"`
	Status       int
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

func (c *Category) TableName() string {
	return TableName("category")
}

func GetCategoryList(page, pageSize int) ([]*Category, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Category, 0)
	query := orm.NewOrm().QueryTable(new(Category).TableName())
	total, _ := query.Count()
	query.OrderBy("order").Limit(pageSize, offset).All(&list)
	return list, total
}

func GetCategoryById(id int64) (Category, error) {
	var list Category
	query := orm.NewOrm().QueryTable(new(Category).TableName())
	query.Filter("id", id).Filter("status", 0).One(&list)
	return list, nil
}

func CategoryAdd(c *Category) (int64, error) {
	return orm.NewOrm().Insert(c)
}

func (c *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}
