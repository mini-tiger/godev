package main

import (
	"tjtools/db/redis"
	"log"
	"fmt"
)

var Redis111 redis.RedisS

func main() {
	client1, err := redis.CreateClient(8, "192.168.43.11:6379", "")
	if err != nil {
		log.Println("NewRedis err:", err)
	}
	defer client1.Close()
	Redis111.Conn = client1

	go func() {
		for {
			s,err:=Redis111.SubChan("abc")
			if err!=nil{
				log.Println("sub err",err)
				continue
			}
			fmt.Println("receive,",)
		}


	}()

	select {}
}
