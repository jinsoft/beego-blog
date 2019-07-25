package admin

import (
	"html/template"
	"strings"
	"time"
	"web/models"
	"web/util"
)

type UserController struct {
	baseController
}

// 用户列表
func (c *UserController) Index() {
	if c.IsAjax() {
		page, err := c.GetInt("page")
		if err != nil {
			page = 1
		}
		limit, err := c.GetInt("limit")
		if err != nil {
			limit = 15
		}
		name := strings.TrimSpace(c.GetString("name"))
		// 查询条件
		filters := make([]interface{}, 0)

		if name != "" {
			filters = append(filters, "name_icontains", name)
		}
		result, count := models.GetUserList(page, limit, filters...)
		list := make([]map[string]interface{}, len(result))
		for k, v := range result {
			row := make(map[string]interface{})
			row["id"] = v.Id
			row["uid"] = v.Uid
			row["name"] = v.Name
			row["email"] = v.Email
			row["phone"] = v.Phone
			if v.LastLogin.IsZero() {
				row["last_login"] = "-"
			} else {
				row["last_login"] = v.LastLogin.Format("2006-01-02 15:04")
			}
			if v.Avatar == "" {
				v.Avatar = "/static/img/default.jpg"
			}
			row["avatar"] = v.Avatar
			row["forbidden"] = v.Forbidden
			row["created_time"] = v.CreatedTime.Format("2006-01-02 15:04")
			list[k] = row
		}
		c.Data["json"] = map[string]interface{}{
			"code":  0,
			"msg":   "请求成功",
			"count": count,
			"data":  list,
		}
		c.ServeJSON()
		c.StopRun()
	}
	//var page int64
	//var pagesize int64 = 15
	//var list []*models.User
	//
	//if page, _ = c.GetInt64("page"); page < 1 {
	//	page = 1
	//}
	//
	//offset := (page - 1) * pagesize
	//
	//count, _ := user.Query().Count()
	c.display()
}

func (c *UserController) Create() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.IsAjax() {
		User := new(models.Users)
		User.Uid = util.GUID()
		User.Name = strings.TrimSpace(c.GetString("name"))
		User.Phone = strings.TrimSpace(c.GetString("phone"))
		User.Email = strings.TrimSpace(c.GetString("email"))
		User.Avatar = strings.TrimSpace(c.GetString("avatar"))
		//User.Password =strings.TrimSpace(c.GetString("password"))
		_, err := models.GetUserByEmail(User.Email)

		if err == nil {
			c.ajaxMsg("邮箱已存在", MSG_ERR)
		}
		pwd := strings.TrimSpace(c.GetString("passwrod"))
		User.Password = models.Md5([]byte(pwd))
		User.CreatedTime = time.Now()
		if _, err := models.UserAdd(User); err != nil {
			c.ajaxMsg(err.Error(), MSG_ERR)
		}
		c.ajaxMsg("添加成功", MSG_OK)
	}
	c.display()
}

func (c UserController) Edit() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	//uid := c.Ctx.Input.Param(":uid")
	//uid := "521c73534fcc1a1636dbf8cd3e2c1d5b"
	//if c.IsAjax() {
	//	User := new(models.Users)
	//	User.Name = strings.TrimSpace(c.GetString("name"))
	//}
	//User,_ :=models.GetUserByUid(uid)
	//row := make(map[string]interface{})
	//row["uid"] = uid
	//row["name"] = User.Name
	//row["email"] = User.Email
	//row["phone"] = User.Phone
	//row["avatar"] = User.Avatar
	//c.Data["user"] = row
	c.TplName = "admin/user/edit.html"
}

func (c *UserController) Table() {
	c.Data["code"] = 0
	c.Data["msg"] = "请求成功"
	c.Data["count"] = 100
	c.ServeJSON()
}
