package models

import (
	"github.com/astaxie/beego/orm"
)

func RegitDB(dbstr string) {
	orm.RegisterModel(new(EventCases))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)

	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql",
		dbstr,
		maxIdle, maxConn)
}
