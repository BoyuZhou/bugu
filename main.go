package main

import (
	"fmt"

	_ "bugu/routers"
	"bugu/redis"
	"bugu/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	utils.InitLogs();

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		beego.AppConfig.String("DataBaseUser"),
		beego.AppConfig.String("DataBasePassword"),
		beego.AppConfig.String("DataBaseHost"),
		beego.AppConfig.String("DataBasePort"),
		beego.AppConfig.String("DataBaseName"))
	orm.RegisterDataBase("default", "mysql", dataSource+"?charset=utf8&loc=Local")

	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 100)

	orm.Debug = beego.BConfig.RunMode == "dev"
	redis.RegisterClient()
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
