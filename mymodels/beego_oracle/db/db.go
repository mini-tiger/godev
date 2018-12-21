package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"godev/mymodels/beego_oracle/models"
)

func Init() {
	// todo 注册DB
	models.RegitDB("test1/test1@192.168.43.236:1521/orcl")
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
