package main

import (
	_ "github.com/go-sql-driver/mysql"
	"godev/mymodels/beego_models/db"
	"godev/mymodels/beego_models/utils"
)

func main() {
	db.Init()
	utils.CRUD("v2", "agent", "0edb4518-1fe2-476b-bd3c-f38e1a81b821")

}
