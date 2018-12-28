package main

import (
	"fmt"
	"tjtools/db/redis"
	"log"
	"time"
)

var Redis1 redis.RedisS
type VehicleJson struct {
	TRAIN_ID       string   `json:"TRAIN_ID"`
	PASS_TIME_CHAR string   `json:"PASS_TIME"`
	INDEX_ID       string   `json:"FOLDER"`
	VehicleOrder   string   `json:"VEHICLE_ORDER"`
	VehicleType    string   `json:"VEHICLE_TYPE"`
	POS_AB_FALG    string   `json:"AB"`
	Images         []string `json:"Images"`
	SaveUnixStamp  int      `json:"Save_unix_stamp,omitempty"` // todo 发送时要 清空此项, 存入redis对比 时间是否超过最大    "maxvehicletime":120
	Send           int     `json:"Send,omitempty"`            // todo 是否发送过 ，存入redis对比 时间是否超过最大
	Vehicle_serial string   `json:"Vehicle_serial,omitempty"`            // todo 程序需要使用，发送前清空
}
func initclient() {
	client1, err := redis.CreateClient(8,"192.168.43.11:6379", "")
	if err != nil {
		log.Println("NewRedis err:", err)
	}
	defer client1.Close()

	Redis1.Conn = client1

	//Redis1.StingJson() // 存储json

	//Redis1.StringOperation() // 存入 获取字符串
	//fmt.Println("===============================")
	//Redis1.StringSet("a",11)
	//Redis1.StringSet("a",22)
	var v VehicleJson
	v.Send=0
	//d :=false
	//v.Send = d
	v.Vehicle_serial = "06020E2995A743F09CE1B0F52F174FF0"

	//fmt.Println(Redis1.StringExists("06020E2995A743F09CE1B0F52F174FF0")) // KEY是否存在

	Redis1.SetJson("06020E2995A743F09CE1B0F52F174FF0",&v,time.Duration(0)*time.Second)

	v1:=VehicleJson{}
	Redis1.GetJson("06020E2995A743F09CE1B0F52F174FF0",&v1)
	fmt.Printf("%+v\n",v1.Send)

	//val,err:=Redis1.StringGet("06020E2995A743F09CE1B0F52F174FF0") // K
	//if val==""{
	//	fmt.Println("key nil")
	//}
	//fmt.Println(val)
	//
	//
	//fmt.Println("===============================")
	//Redis1.listOperation() // list操作
	//fmt.Println("===============================")

	//Redis1.Set() // set操作
	//redis.SetOperation(client1)

	fmt.Println("===============================")

	//Redis1.pubsub1() // 订阅 消息分发

	//setOperation(client)
	//hashOperation(client)
	//
	//connectPool(client)
	select {}
}
func main() {

	initclient()
}
