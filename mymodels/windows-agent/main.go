package main

import (
	"flag"
	"fmt"
	"godev/mymodels/windows-agent/cron"
	extend_cron "godev/mymodels/windows-agent/extend/cron"
	"godev/mymodels/windows-agent/funcs"
	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/http"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitLog()

	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()
	//
	funcs.BuildMappers()
	//
	go cron.InitDataHistory()
	//
	cron.ReportAgentStatus()
	//
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()

	cron.LoadEnvironmentGridConfig()

	extend_cron.Loadportporcess_taskConfig()
	go extend_cron.Updateportprocess_env_task()

	//cron.UploadEnvironmentGrid() //硬件信息

	cron.Collect()
	//
	go http.Start()

	select {}
}
