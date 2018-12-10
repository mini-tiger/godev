package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"godev/mymodels/falconApi/models"
	_ "godev/mymodels/falconApi/routers"
)

func main() {
	models.Init()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
