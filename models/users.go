package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	Id          int64
	Uid         string
	Email       string
	Password    string
	Name        string
	Phone       string
	Forbidden   int
	LastLogin   time.Time
	CreatedTime time.Time
	//Logincount  int64
	//Lastloginip string
	Avatar string
}

func (m *Users) TableName() string {
	return TableName("users")
}

/*func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
*/

func GetUserList(page, pageSize int, filters ...interface{}) ([]*Users, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Users, 0)
	query := orm.NewOrm().QueryTable(new(Users).TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func GetUserByEmail(email string) (*Users, error) {
	u := new(Users)
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("email", email).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserAdd(u *Users) (int64, error) {
	return orm.NewOrm().Insert(u)
}
