package main

import (
	"tjtools/db/redis"
	"log"
	"time"
)

var Redis11 redis.RedisS

func main() {
	client1, err := redis.CreateClient(8, "192.168.43.11:6379", "")
	if err != nil {
		log.Println("NewRedis err:", err)
	}
	defer client1.Close()
	Redis11.Conn = client1
	go func() { //生产者
		for{
			err:=Redis11.PubChan("abc","111")
			if err!=nil{
				log.Println("pub err:",err)
			}

			time.Sleep(time.Duration(2)*time.Second)
		}
	}()


	select {}
}
