package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db_host := beego.AppConfig.String("DB_HOST")
	db_port := beego.AppConfig.String("DB_PORT")
	db_databases := beego.AppConfig.String("DB_DATABASE")
	db_username := beego.AppConfig.String("DB_USERNAME")
	db_password := beego.AppConfig.String("DB_PASSWORD")

	if db_port == "" {
		db_port = "3306"
	}

	// username:password@protocol(address)/dbname?param=value
	db_url := db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_databases + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", db_url)

	orm.RegisterModel(new(Users), new(Admins), new(Article), new(Category))
	//orm.RegisterModelWithPrefix(beego.AppConfig.String("DB_PREFIX"),
	//	new(Users), new(Admins))

	// 开发环境开启debug
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("DB_PREFIX"), str)
}
