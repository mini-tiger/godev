package main

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
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
func (p *program) run() {


	//f, err := os.OpenFile("c:\\abc.log", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755) //读写模式
	//if err != nil {
	//	os.Exit(1)
	//}
	//defer func() {
	//	f.Close()
	//}()
	os.Chdir("C:\\GclAgentWin")
	g.InitRootDir()

	g.ParseConfig("")


	//g.InitLog(f)

	g.InitLogging()

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
	<-time.After(time.Second * 5)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GclMonitor",
		DisplayName: "GclMonitor",
		Description: "GclMonitor",
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