package main

import (
	"godev/mymodels/gorm/utils"
	"godev/mymodels/gorm/models"
	"fmt"
)



func main()  {
	utils.Init_db()
	db:=utils.Conn().Uic

	var u models.User
	db.Where("name=?","admin").Find(&u)
	fmt.Println(u)

	}

