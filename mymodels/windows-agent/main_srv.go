// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"log"

	"fmt"
	"github.com/kardianos/service"
	"os"
	"godev/mymodels/windows-agent/cron"
	"godev/mymodels/windows-agent/funcs"
	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/http"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() { // todo 可以使用exec.command() 加入写过的程序


	g.InitRootDir()
	g.ParseConfig()
	g.InitLog()


	g.InitLocalIps()
	g.InitRpcClients()

	cron.CollectInfo()
	//
	g.RunStatus()
	funcs.BuildMappers()
	//
	go cron.InitDataHistory()
	//
	//cron.ReportAgentStatus() // 传送硬件信息到数据库，windows版本不支持plugin
	//
	g.LoadUUIDBIZ()
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()

	cron.LoadEnvironmentGridConfig() // agent 机器的删除尽量在，前端界面，涉及多表
	// 直接删除库后， 再次运行Agent，不往graph库endpoint中插入,要删除graph/data/6070文件夹
	//extend_cron.Loadportporcess_taskConfig()
	//extend_cron.Updateportprocess_env_task()

	cron.UploadEnvironmentGrid() //硬件信息

	cron.Collect()
	//
	go http.Start()

	select {}


}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func (p *program) Install(s service.Service) error {
	err := s.Install()
	if err != nil {
		fmt.Printf("Failed to install: %s\n", err)
		return err
	}
	fmt.Printf("Service \"%s\" installed.\n", "install system server")
	return nil
}

func (p *program) Uninstall(s service.Service) error {
	err := s.Uninstall()
	if err != nil {
		fmt.Printf("Failed to install: %s\n", err)
		return err
	}
	fmt.Printf("Service \"%s\" installed.\n", "install system server")
	return nil
}

var s service.Service

func main() {
	svcConfig := &service.Config{
		Name:        "GCL Monitor Windows Agent",
		DisplayName: "GCL Monitor Windows Agent",
		Description: "GCL Monitor Windows Agent",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Printf("service.new: %s\n", err)
	}
	//s.Uninstall()
	//s.Install()
	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err := s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
			}
		case "remove":
			err = s.Uninstall()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", "go")
		case "run":
			s.Run()
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", "")
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", "go")
		}
		return
	}

	if err != nil {
		log.Fatal(err)
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
