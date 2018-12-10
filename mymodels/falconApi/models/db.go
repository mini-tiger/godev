package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	// todo 注册DB
	RegitDB("falcon:123456@tcp(192.168.1.104:3306)/alarms?charset=utf8&loc=Asia%2FShanghai")
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
