package main

import (
	"tjtools/db/redis"
	"log"
	"fmt"
	"encoding/json"
)

var Redis111 redis.RedisS

type RedisJson struct {
	Name string
	Num  int
}

func main() {
	client1, err := redis.CreateClient(8, "192.168.43.11:6379", "")
	if err != nil {
		log.Println("NewRedis err:", err)
	}
	defer client1.Close()
	Redis111.Conn = client1

	go func() {
		for {
			s, err := Redis111.SubChan("abc")
			if err != nil {
				log.Println("sub err", err)
				continue
			}
			fmt.Println("receive,", s.Pattern,s.Channel,s.Payload)
			var rj RedisJson
			err = json.Unmarshal([]byte(s.Payload), &rj)
			if err != nil {
				log.Println("sub json err", err)
			}
			fmt.Printf("解析json后： %+v\n",rj)
		}

	}()

	select {}
}
