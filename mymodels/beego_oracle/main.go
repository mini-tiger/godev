package main

import (
	_ "github.com/mattn/go-oci8"
	"godev/mymodels/beego_oracle/db"
)

func main() {
	db.Init()
	//utils.CRUD("v2", "agent", "0edb4518-1fe2-476b-bd3c-f38e1a81b821")

}
