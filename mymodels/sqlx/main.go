package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//const (
//	INTERVAL = 10
//)

type Hostinfo struct {
	ID         int    `json:"id" `
	Uuid       string `json:"uuid"`
	Ip         string `json:"ip"`
	Passwd     string `json:"-"`
	Username   string `json:"username"`
	Port       int    `json:"port"`
	Bizid      int    `json:"bizid" gorm:"column:bizid"`
	Createtime string `json:"createtime"`
}

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "falcon:123456@tcp(192.168.43.11:3306)/nodeman")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}

func main() {

	s := selectEx("hostinfo")  //todo 不同表统一返回
	fmt.Printf("%T,%+v\n", s, s)
}

func selectEx(dbname string) (ret interface{}) {
	//var p interface{}
	//var ret interface{}
	switch dbname {
	case "hostinfo":
		p := make([]Hostinfo, 0)
		err := Db.Select(&p, "select ip,port from hostinfo where id=?", 1)

		if err != nil {
			fmt.Println("exec failed, ", err)
			//return
		}
		ret = p
	}
	return ret

}
