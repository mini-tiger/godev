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


type Endpoint struct {
	ID               uint              `gorm:"primary_key"`
	Endpoint         string            `json:"endpoint"`
	Ts               int               `json:"-"`
	TCreate          time.Time         `json:"-"`
	TModify          time.Time         `json:"-"`
	EndpointCounters []EndpointCounter `gorm:"ForeignKey:EndpointIDE"`
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

	eids:=fmt.Sprintf("(%d.%d)",517,141) // 需要带括号的格式 in
	dt := db.Table("endpoint_counter").Select("endpoint_id, counter, step, type").Where(fmt.Sprintf("endpoint_id IN %s", eids))
	// Select where 级联调用
	dt = dt.Where("counter regexp ?", strings.TrimSpace("df"))
	//todo 将结果再次where，limit 分页用  ， scan将结果存入 counteres切片，每个元素按照Struct存储
	dt = dt.Limit(50).Offset(0).Scan(&counters)
	for _,v:=range counters  {
		fmt.Printf("%+v\n",v)
	}

	fmt.Printf("%s\n","====================================")

	type DBRows struct {
		Endpoint  string
		CounterId int
		Counter   string
		Type      string
		Step      int
	}
	inputs:=[]string{}
	inputs=append(inputs, []string{"falcon-win12-1","localhost.localdomain"}... )
	rows := []DBRows{}
	dt = db.Raw(
		`select a.endpoint, b.id AS counter_id, b.counter, b.type, b.step from endpoint as a, endpoint_counter as b
		where b.endpoint_id = a.id
		AND a.endpoint in (?)`, inputs)

	dt.Limit(10).Scan(&rows)
	for _,v:=range rows  {
		fmt.Printf("%+v\n",v)
	}


	fmt.Printf("%s\n","====================================")

	//del_r := Endpoint{}
	tx := db.Begin()
	dt = tx.Table("endpoint").Where("endpoint in (?)", inputs).Delete(&Endpoint{})
	//fmt.Println(Endpoint{})
	fmt.Println(dt.RowsAffected)
	}


