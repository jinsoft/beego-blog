package admin

import "web/models"

type UserController struct {
	baseController
}

// 用户列表
func (c *UserController) List() {
	var page int64
	var pagesize int64 = 15
	var list []*models.User

	if page, _ = c.GetInt64("page"); page < 1 {
		page = 1
	}

	offset := (page - 1) * pagesize

	count, _ := user.Query().Count()
}
