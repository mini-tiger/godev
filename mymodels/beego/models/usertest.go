package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

func RegitDB() {
	orm.RegisterModel(new(User), new(Class))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)

	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql",
		"falcon:123456@tcp(192.168.43.11:3306)/test?charset=utf8&loc=Asia%2FShanghai",
		maxIdle, maxConn)
}

type User struct {
	Id      int       `orm:"pk;auto"`
	Name    string    `orm:"index;column(username);size(60)"`
	Update  time.Time `orm:"auto_now_add;type(datetime)"`
	Email   string    `orm:"column(email);size(60)"`
	Class []*Class  `orm:"reverse(many)"`
}

type Class struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"size(20)"`
	User *User  `orm:"rel(fk);on_delete(cascade)"`
}

//自定义表名
func (u *User) TableName() string {
	return "auth_user"
}

// 多字段索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Name"},
	}
}

// 多字段唯一键
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Email"},
	}
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}
