package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int64
	Uid         string
	Email       string
	Password    string
	Nickname    string
	Lastlogin   time.Time
	Logincount  int64
	Lastloginip string
	Avator      string
}

func (m *User) TableName() string {
	return TableName("user")
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
