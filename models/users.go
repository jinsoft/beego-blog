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
	Status      int
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

func GetUserByUid(uid string) (*Users, error) {
	u := new(Users)
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("uid", uid).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserById(id int64) (*Users, error) {
	u := new(Users)
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Add(u *Users) (int64, error) {
	return orm.NewOrm().Insert(u)
}

func (c *Users) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (u *Users) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}
