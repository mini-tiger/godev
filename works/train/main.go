package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godev/works/train/g"
	"godev/works/train/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")

	flag.Parse()

	g.ParseConfig(*cfg)

	g.InitRedis()
	g.InitLog1()

	db.Init()
	/*
		todo  任务下发到update后，如果任务没有执行完成前update请求任务，会重复发送，已问题可以 使用redis存储任务执行状态
		redis 已经写了 set组合操作
	*/

	//cache.Init()

	//go cache.DeleteStaleAgents()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		//db.DB.Close()
		os.Exit(0)
	}()

	select {}
}
