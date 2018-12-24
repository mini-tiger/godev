package main

import (
	"fmt"
	"godev/tjtools/db/redis"
	"log"
)

var Redis1 redis.RedisStruct

func initclient() {
	client1, err := redis.NewRedisClient("192.168.43.11:6379", "", 0, 10)
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
	Redis1.StringExists("a") // KEY是否存在
	val,err:=Redis1.StringGet("a") // K
	if val==""{
		fmt.Println("key nil")
	}
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
