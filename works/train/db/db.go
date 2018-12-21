package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/open-falcon/falcon-plus/modules/ubs/g"
	"github.com/open-falcon/falcon-plus/modules/ubs/models"
	"log"
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

func Init() {
	// todo ×¢²áDB
	models.RegitDB(g.Config().Database)
	// Êý¾Ý¿â±ðÃû
	name := "default"
	// drop table ºóÔÙ½¨±í
	force := false
	// ´òÓ¡Ö´ÐÐ¹ý³Ì
	verbose := true
	// Óöµ½´íÎóÁ¢¼´·µ»Ø
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		log.Fatalf("db %s err:\n", g.Config().Database, err)
	}
	orm.Debug = false

}
