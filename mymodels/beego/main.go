package main

import (
	_ "godev/mymodels/beego/routers"
	"github.com/astaxie/beego"
	"godev/mymodels/beego/models"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func main() {
	// todo 注册DB
	models.RegitDB()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.Debug = true

	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.Run()
}

