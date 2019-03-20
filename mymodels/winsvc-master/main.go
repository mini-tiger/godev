package main

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)
/*
注意代码中 服务名的定义
 打包
go build -ldflags="-s -w -H=windowsgui -linkmode external -extldflags '-static'" -x -o test.exe main.go

安装成服务
test.exe install

在services.msc中启动服务
todo 服务方式启动默认在 c:\windows\system32下工作
路径拼接使用filepath.join() 不要 path.join()
日志打印使用 windows-agent\g\log_diy ，nx4log包不能使用

*/

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	f, err := os.OpenFile("c:\\abc.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755) //读写模式
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		f.Close()
	}()

	begintime := time.Now().String()
	f.WriteString(begintime)
	f.Write([]byte("\n"))
	go func() {
		for {
			str1 := "aaaaa,bbbbbbbb,cccccccccc\n"
			ws := []byte(str1)
			f.Write(ws) //追加写
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
	select {}

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	<-time.After(time.Second * 5)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleStopPause",
		DisplayName: "Go Service Example: Stop Pause",
		Description: "This is an example Go service that pauses on stop.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
