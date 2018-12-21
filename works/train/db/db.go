package db

import (
	"godev/tjtools/db/oracle"
	"godev/works/train/g"
	"reflect"
	"fmt"
)
//var DB *sql.DB
//
//func Init() {
//	var err error
//	DB, err = sql.Open("mysql", g.Config().Database)
//	if err != nil {
//		log.Fatalln("open db fail:", err)
//	}
//
//	DB.SetMaxOpenConns(g.Config().MaxConns)
//	DB.SetMaxIdleConns(g.Config().MaxIdle)
//
//	err = DB.Ping()
//	if err != nil {
//		log.Fatalln("ping db fail:", err)
//	}
//}


var DB oracle.OrclData

func Init() {
	err,db:=oracle.NewOrclClient("test1/test1@1.119.132.155:1521/orcl")
	if err!=nil{
		g.Logger().Error("db init Fail")
		return
	}
	g.Logger().Printf("db success")
	DB = oracle.OrclData{DB:db}

}

func GetRows(sql string,i interface{},structName string) (err error,data []interface{}) {
	//typ:=reflect.TypeOf(i)
	v:=reflect.ValueOf(i)
	//fmt.Println(typ.Kind() == reflect.Ptr,typ.Name())
	fmt.Println(v.Type(),v.Kind() == reflect.Ptr)

	err,rows:=DB.GetSqlData(sql)
	if err!=nil{
		g.Logger().Error("get sql %s err:%s",sql,err)
	}

	defer rows.Close()

	for rows.Next() {
		//var data string

		rows.Scan(&data)
		//fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		return
	}
}


