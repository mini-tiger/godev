package main

import (
	//log "github.com/sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
	"time"
)

type Animal struct {
	Name string
	age  int
}

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	a := Animal{"dog", 22}
	log.SetLevel(log.DebugLevel) //级别
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})
	//log.SetOutput()  todo 参考log 同时写入文件 OS.STDOUT
	logcustom := log.WithFields(log.Fields{ //自定义格式
		"event": "ne",
		"topic": "title",
		"key": "my key",
	})

	log.Error("hello world")
	for {
		time.Sleep(time.Second)
		a.age++
		logcustom.Info(a)
		log.Printf("i am ok %s", "dock")
	}
	log.Fatal("kill ")
}
