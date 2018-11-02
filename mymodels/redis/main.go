package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"log"
)

type RedisStruct struct {
	Conn redis.Conn
	Lock sync.RWMutex
}

var Redis RedisStruct

func simpleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (r RedisStruct) set(k, v string) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	_, err := r.Conn.Do("set", k, v)
	simpleErr(err)
	return err
}

func (r RedisStruct) get(k string) (error, string) {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	value, err := redis.String(r.Conn.Do("GET", k))
	simpleErr(err)
	if value == ""{
		fmt.Println("nil")
	}
	fmt.Println(err,value)
	return err, value
}

func initredis() {
	c, err := redis.Dial("tcp", "192.168.43.11:6378")
	if err != nil {
		fmt.Println("Connect to redis error", err)

	}
	Redis.Conn = c
	//defer c.Close()

}

func main() {
	initredis()
	//Redis.set("a", "1")
	Redis.get("b")

}
