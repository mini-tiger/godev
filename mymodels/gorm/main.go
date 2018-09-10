package main

import (
	"godev/mymodels/gorm/utils"
	"godev/mymodels/gorm/models"
	"fmt"
	"strings"
	"time"
)

type EndpointCounter struct {
	ID         uint `gorm:"primary_key"`
	EndpointID int
	Counter    string
	Step       int
	Type       string
	Ts         int
	TCreate    time.Time
	TModify    time.Time
}


func main()  {
	utils.Init_db()
	db:=utils.Conn().Uic

	var u models.User
	db.Where("name=?","admin").Find(&u)
	fmt.Println(u)

	fmt.Printf("%s\n","====================================")
	db=utils.Conn().Graph

	var counters []EndpointCounter

	eids:=fmt.Sprintf("(%d.%d)",517,141)
	dt := db.Table("endpoint_counter").Select("endpoint_id, counter, step, type").Where(fmt.Sprintf("endpoint_id IN %s", eids))
	dt = dt.Where("counter regexp ?", strings.TrimSpace("df"))
	dt = dt.Limit(50).Offset(0).Scan(&counters)
	for _,v:=range counters  {
		fmt.Printf("%+v\n",v)
	}

	fmt.Printf("%s\n","====================================")


	}

