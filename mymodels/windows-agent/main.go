package main

import (
	"flag"
	"fmt"
	"godev/mymodels/windows-agent/cron"
	extend_cron "godev/mymodels/windows-agent/extend/cron"
	"godev/mymodels/windows-agent/funcs"
	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/http"
	"log"
	"net"
	httpex "net/http"
	"os"
	"strings"
	"time"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()
	go func() {
		conn, err := net.DialTimeout("tcp", "www.baidu.com:80", 2*time.Second)
		if err != nil {
			log.Println("net.DialTimeout Fail")
			return
		}
		ip := conn.LocalAddr().String()
		ips := strings.Split(ip, ":")
		log.Println(httpex.ListenAndServe(fmt.Sprintf("%s:7777", ips[0]), nil))
		log.Println(ips[0], "7777 port listen sucess")
	}()
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
	cron.ReportAgentStatus() // 传送硬件信息到数据库，windows版本不支持plugin
	//
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()

	cron.LoadEnvironmentGridConfig() // agent 机器的删除尽量在，前端界面，涉及多表
	// 直接删除库后， 再次运行Agent，不往graph库endpoint中插入,要删除graph/data/6070文件夹
	extend_cron.Loadportporcess_taskConfig()
	extend_cron.Updateportprocess_env_task()

	cron.UploadEnvironmentGrid() //硬件信息

	cron.Collect()
	//
	go http.Start()

	select {}
}
