package main

import (
	"tjtools/db/redis"
	"log"
	"time"
	"encoding/json"

)

var Redis11 redis.RedisS


type RedisJson1 struct {
	Name string
	Num int
}
func main() {
	client1, err := redis.CreateClient(8, "192.168.43.11:6379", "")
	if err != nil {
		log.Println("NewRedis err:", err)
	}
	defer client1.Close()
	Redis11.Conn = client1
	go func() { //生产者
		for i := 0; ; i++ {
			var rj RedisJson1
			rj.Name = "hh"
			rj.Num = i
			bj, err := json.Marshal(rj)

			if err != nil {
				log.Println("pub json err:", err)
			}

			err = Redis11.PubChan("abc", string(bj))
			if err != nil {
				log.Println("pub err:", err)
			}

			time.Sleep(time.Duration(2) * time.Second)
		}
	}()

	select {}
}
