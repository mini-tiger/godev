package models

import (
	"github.com/astaxie/beego/orm"
)

func RegitDB(dbstr string) {
	orm.RegisterModel(new(Mission))
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

type Mission struct {
	Id          int    `orm:"pk;auto"`
	AppName     string `orm:"index;column(appname);size(20)"` // 应用名字
	Version     string `orm:"column(version);size(20)"`       //版本
	UUID        string `orm:"column(uuid);size(60);index"`    // agent uuid
	Status      uint   `orm:"column(status);default(0)"`      // 是否安装过   0 没有 1 有
	Count       uint   `orm:"column(count);default(0)"`       // 安装过的次数
	Success     bool   `orm:"column(success);default(false)"` // 安装是否成功
	CreateTime  int64  `orm:"column(createtime)"`             // 创建时间 unix time
	LastTime    int64  `orm:"column(lasttime)"`               // 最后一次安装时间
	InstallTime int64  `orm:"column(installtime)"`            // 计划安装时间
}

//自定义表名
func (m *Mission) TableName() string {
	return "mission"
}

// 多字段唯一键
func (u *Mission) TableUnique() [][]string {
	return [][]string{
		[]string{"AppName", "Version", "UUID"},
	}
}

// 设置引擎为 INNODB
func (u *Mission) TableEngine() string {
	return "INNODB"
}
