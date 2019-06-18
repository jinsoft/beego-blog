package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	db_url := db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_databases + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", db_url)

	orm.RegisterModel(new(User))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
