package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id       int64
	Uid      string `org:"size(32)"`
	Name     string `org:"size(128)"`
	Email    string `json:"email" orm:"unique;size(128);index"`
	Password string `orm:"size(64)"`
}

func (m *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Admin) TableName() string {
	return TableName("admins")
}
