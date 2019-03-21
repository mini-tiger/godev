package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"godev/mymodels/beego_models/models"
)

func Init() {
	// todo 注册DB
	models.RegitDB("root:W3b5Ev!c3@tcp(bi.itma.com.cn:3306)/bi_yunji?charset=utf8&loc=Asia%2FShanghai")
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println("err:",err)
	}
	orm.Debug = false
}
