package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"godev/mymodels/beego_models/models"
)

func Init() {
	// todo 注册DB
	models.RegitDB("falcon:123456@tcp(192.168.43.11:3306)/app?charset=utf8&loc=Asia%2FShanghai")
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
	orm.Debug = false
}
