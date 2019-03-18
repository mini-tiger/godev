package g

import (
	"bytes"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	nxlog "github.com/ccpaging/nxlog4go"
	"godev/mymodels/windows-agent/common/model"
	"github.com/toolkits/net"
	"github.com/toolkits/slice"
	"io"

	"path"
)

var (
	Root  string
	logge *nxlog.Logger
	OutIP *string
)

type Log1 struct{}

func (l *Log1) Println(m ...interface{}) {
	logge.Info(m)
}
func (l *Log1) Printf(arg0 interface{}, args ...interface{}) {
	//fmt.Printf("%T,%v\n", m, m)
	//arg0 = arg0.(string)

	//fmt.Printf("%T,%v\n",arg0,arg0)
	arg0 = strings.Trim(arg0.(string),"\n") // todo 去掉换行
	logge.Info(arg0, args...)
}
func (l *Log1) Error(m ...interface{}) {
	logge.Error(m)
}

func (l *Log1) Debug(arg0 interface{}, args ...interface{}) {
	logge.Debug(arg0 , args...)
}

func (l *Log1) Fatalln(m ...interface{}) {
	logge.Error(m)
	os.Exit(1)
}
func (l *Log1) Fatalf(arg0 interface{}, args ...interface{}) {
	logge.Error(arg0, args)
	os.Exit(1)
}

var logger *Log1

func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}
//type Logstruct struct {
//	sync.RWMutex
//	Log *log.Logger
//}
//
//func (l *Logstruct)Info(str string)  {
//	l.Log.SetPrefix("[INFO]")
//	l.Log.Print(str)
//}
//
//func (l *Logstruct)Error(str string)  {
//	l.Log.SetPrefix("[ERROR]")
//	l.Log.Print(str)
//}
//
//var log11 Logstruct

func InitLog() *nxlog.Logger{
	//fileName := Config().Logfile

	//logFile, err := os.Create(fileName)
	//if err != nil {
	//	log.Fatalln("open file error !")
	//}
	fileName:=path.Join(Root,Config().Logfile)


	//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatalln("fail to create test.log file!", err)
	//}
	//log11.Log = log.New(file, "", log.Ltime|log.Ldate)
	//log.Println("1.Println log with log.LstdFlags ...")
	//logger.Println("1.Println log with log.LstdFlags ...")


	//go func() {
	//	for{
	//		log11.Info(fmt.Sprintln(1111))
	//		time.Sleep(time.Duration(1)*time.Second)
	//	}
	//}()

	nxlog.FileFlushDefault = 5 // 修改默认写入硬盘时间
	nxlog.LogCallerDepth = 3 //runtime.caller(3)  日志触发上报的层级
	rfw := nxlog.NewRotateFileWriter(fileName).SetDaily(true).SetMaxBackup(Config().LogMaxDays) //log保存最大天数



	var ww io.Writer
	if Config().Daemon{
		ww = io.MultiWriter(rfw) //todo 输出到rfw定义
	}else{
		ww = io.MultiWriter(os.Stdout,rfw) //todo 同时输出到rfw 与 系统输出
	}

	// Get a new logger instance
	// todo FINEST 级别最低
	// todo %p prefix, %N 行号
	logge = nxlog.New(nxlog.FINEST).SetOutput(ww).SetPattern("%P [%Y %T] [%L] (%S LineNo:%N) %M\n")
	//Log.SetPrefix("11111")
	logge.SetLevel(1)

	logge.Info("read config file ,successfully") // 走到这里代表配置文件已经读取成功
	logge.Info("日志文件最多保存%d天",Config().LogMaxDays)
	logge.Info("logging on %s", fileName)
	logge.Info("进程已启动, 当前进程PID:%d",os.Getpid())
	return logge
	// Log some experimental messages
	//for j := 0; j < 15; j++ {
	//	for i := 0; i < 400 / (j+1); i++ {
	//		Log.Finest("Everything is created now (notice that I will not be printing to the file)")
	//		Log.Info("%d. The time is now: %s", j, time.Now().Format("15:04:05 MST 2006/01/02"))
	//		Log.Critical("Time to close out!")
	//
	//		time.Sleep(1*time.Second)
	//	}
	//}
	//rfw.Close()

	//logger = log.New(logFile, "[Debug]", log.LstdFlags)

}

func Logger() *Log1 {
	lock.RLock()
	defer lock.RUnlock()
	return logger
}

var LocalIps []string

func InitLocalIps() {
	var err error
	LocalIps, err = net.IntranetIP()
	if err != nil {
		//logger.Fatalln("get intranet ip fail:", err)
		logger.Error("get intranet ip fail:", err)
		os.Exit(1)
	}
}

var (
	HbsClient *SingleConnRpcClient
	UbsClient *SingleConnRpcClient
)

func InitRpcClients() {
	if Config().Heartbeat.Enabled {
		HbsClient = &SingleConnRpcClient{
			RpcServer: Config().Heartbeat.Addr,
			Timeout:   time.Duration(Config().Heartbeat.Timeout) * time.Millisecond,
		}
	}
	if Config().Ubs.Enabled {
		UbsClient = &SingleConnRpcClient{
			RpcServer: Config().Ubs.Addr,
			Timeout:   time.Duration(Config().Heartbeat.Timeout) * time.Millisecond,
		}
	}
}


func SendToTransfer(metrics []*model.MetricValue) {
	if len(metrics) == 0 {
		return
	}
	dt := Config().DefaultTags
	if len(dt) > 0 {
		var buf bytes.Buffer
		default_tags_list := []string{}
		for k, v := range dt {
			buf.Reset()
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(v)
			default_tags_list = append(default_tags_list, buf.String())
		}

		default_tags := strings.Join(default_tags_list, ",")
		for i, x := range metrics {
			buf.Reset()
			if x.Tags == "" {
				metrics[i].Tags = default_tags
			} else {
				buf.WriteString(metrics[i].Tags)
				buf.WriteString(",")
				buf.WriteString(default_tags)
				metrics[i].Tags = buf.String()
			}
		}
	}
	debug := Config().Debug

	if debug {
		logger.Printf("=> <Total=%d> %v\n", len(metrics), metrics[0])
		//logger.Debug("=> <Total=%d> %v\n", len(metrics), metrics[0])
	}

	var resp model.TransferResponse
	SendMetrics(metrics, &resp)

	if debug {
		logger.Println("<=", &resp)
		//logger.Debug("<=", &resp)
	}
}

var (
	reportUrls     map[string]string
	reportUrlsLock = new(sync.RWMutex)
)

func ReportUrls() map[string]string {
	reportUrlsLock.RLock()
	defer reportUrlsLock.RUnlock()
	return reportUrls
}

func SetReportUrls(urls map[string]string) {
	reportUrlsLock.RLock()
	defer reportUrlsLock.RUnlock()
	reportUrls = urls
}

var (
	reportPorts     []int64
	reportPortsLock = new(sync.RWMutex)
)

func ReportPorts() []int64 {
	reportPortsLock.RLock()
	defer reportPortsLock.RUnlock()
	return reportPorts
}

func SetReportPorts(ports []int64) {
	reportPortsLock.Lock()
	defer reportPortsLock.Unlock()
	reportPorts = ports
}

var (
	duPaths     []string
	duPathsLock = new(sync.RWMutex)
)

func DuPaths() []string {
	duPathsLock.RLock()
	defer duPathsLock.RUnlock()
	return duPaths
}

func SetDuPaths(paths []string) {
	duPathsLock.Lock()
	defer duPathsLock.Unlock()
	duPaths = paths
}

var (
	// tags => {1=>name, 2=>cmdline}
	// e.g. 'name=falcon-agent'=>{1=>falcon-agent}
	// e.g. 'cmdline=xx'=>{2=>xx}
	reportProcs     map[string]map[int]string
	reportProcsLock = new(sync.RWMutex)
)

func ReportProcs() map[string]map[int]string {
	reportProcsLock.RLock()
	defer reportProcsLock.RUnlock()
	return reportProcs
}

func SetReportProcs(procs map[string]map[int]string) {
	reportProcsLock.Lock()
	defer reportProcsLock.Unlock()
	reportProcs = procs
}

var (
	ips     []string
	ipsLock = new(sync.Mutex)
)

func TrustableIps() []string {
	ipsLock.Lock()
	defer ipsLock.Unlock()
	return ips
}

func SetTrustableIps(ipStr string) {
	arr := strings.Split(ipStr, ",")
	ipsLock.Lock()
	defer ipsLock.Unlock()
	ips = arr
}

func IsTrustable(remoteAddr string) bool {
	ip := remoteAddr
	idx := strings.LastIndex(remoteAddr, ":")
	if idx > 0 {
		ip = remoteAddr[0:idx]
	}

	if ip == "127.0.0.1" {
		return true
	}

	return slice.ContainsString(TrustableIps(), ip)
}
