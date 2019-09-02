package models

import "github.com/astaxie/beego/orm"

type Tags struct {
	Id    int64
	Name  string `orm:"size(100);index"`
	Count int64  `orm:"index"`
}

func (m *Tags) TableName() string {
	return TableName("tags")
}

func (m *Tags) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tags) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tags) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
