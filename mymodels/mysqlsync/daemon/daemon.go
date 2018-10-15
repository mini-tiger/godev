package godaemon

import (
	"flag"
	"fmt"
	"github.com/toolkits/file"
	slog "log"
	"os"
	"os/exec"
	"strings"
)

//var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true or -d true.")
var godaemon = flag.String("d", "false", "run app as a daemon with -d true or -d false.")

//var cfgfile = flag.String("c", "C:\\work\\go-dev\\src\\godev\\mymodels\\mysqlsync\\cfg.json", "configuration file")
var cfgfile = flag.String("c", "cfg.json", "configuration file")

func cfgExist(cfgfile string) {
	f, _ := os.OpenFile("mysqlsync.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) //加载日志前用
	logger := slog.New(f, "", slog.Ltime|slog.Ldate)

	if !file.IsExist(cfgfile) {
		fmt.Printf("arg -c config file: %s,is not existent. \n", cfgfile)
		logger.Fatalf("arg -c config file: %s,is not existent. \n", cfgfile)
		//*cfgfile = "C:\\work\\go-dev\\src\\godev\\mymodels\\mysqlsync\\cfg.json"
	}
	f.Close()
}

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	//if flag.NArg() == 1 { // flag 以外的参数  有1个代表加了 配置文件路径
	//	cmd = exec.Command(os.Args[0])
	//	cmd.Args = append(cmd.Args, flag.Args()...)
	//	cfgExist(flag.Args()[0])
	//}
	//if flag.NArg() == 0 { // flag 以外的参数 没有 使用默认 cfg.json
	//	cmd = exec.Command(os.Args[0])
	//	cmd.Args = append(cmd.Args, "cfg.json")
	//	cfgExist("cfg.json")
	//}

	cfgExist(*cfgfile)
	if strings.Contains(*godaemon, "true") {
		cmd := exec.Command(os.Args[0])
		cmd.Args = append(cmd.Args, *cfgfile)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		*godaemon = "false"
		os.Exit(0)
	}

	//}else{
	//	cmd := exec.Command(os.Args[0])
	//	if flag.NArg() == 1 { // flag 以外的参数  有1个代表加了 配置文件路径
	//		cmd = exec.Command(os.Args[0])
	//		cmd.Args = append(cmd.Args, os.Args[1:]...)
	//	}
	//	if flag.NArg() == 0{   // flag 以外的参数 没有 使用默认 cfg.json
	//		cmd = exec.Command(os.Args[0],)
	//		cmd.Args = append(cmd.Args, "cfg.json")
	//	}
	//	cmd.Start()
	//	fmt.Println(cmd.Args)
	//	os.Exit(0)
	//}
	//os.Exit(0)
}
