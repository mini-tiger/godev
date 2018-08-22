package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

func Conn() DBpool {
	return dbp
}

type DBpool struct {
	Portaldb *gorm.DB
	Uic      *gorm.DB
	//测试MAP
	Dbs map[string]*gorm.DB
}

var dbp DBpool

func Init_db() (err error) {
	dbsp := DBpool{Dbs: make(map[string]*gorm.DB)}

	var p *sql.DB
	portal, err := gorm.Open("mysql", "falcon:123456@tcp(192.168.43.11:3306)/falcon_portal?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return fmt.Errorf("connect to falcon_portal: %s", err.Error())
	}

	portal.LogMode(true) //打开日志
	portal.Dialect().SetDB(p)
	portal.SingularTable(true)
	//// 全局禁用表名复数
	//db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,否则是users 使用`TableName`设置的表名不受影响
	dbp.Portaldb = portal
	dbsp.Dbs["portaldb"] = portal

	var u *sql.DB
	uic, err := gorm.Open("mysql", "falcon:123456@tcp(192.168.43.11:3306)/uic?charset=utf8&parseTime=True&loc=Local")


	if err != nil {
		return fmt.Errorf("connect to falcon_portal: %s", err.Error())
	}
	uic.LogMode(true) //打开日志
	uic.Dialect().SetDB(u)
	uic.SingularTable(true)
	dbp.Uic = uic
	dbsp.Dbs["uic"] = uic

	//uic.First(&models.User,1)
	//fmt.Println(uic.HasTable("user"))

	return
}

func CloseDB() (err error) {
	err = dbp.Portaldb.Close()
	if err != nil {
		return
	}
	err = dbp.Uic.Close()
	if err != nil {
		return
	}
	return
}
