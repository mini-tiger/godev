package main

import (
	"fmt"

	"gopkg.in/redis.v4"
	"time"
)

type Host struct {
	Ip string
}


func main()  {
	//rc:=createClient()
	fmt.Printf("%T\n",time.Now().Unix())
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.43.11:6378",
		Password: "",
		DB:       0,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
