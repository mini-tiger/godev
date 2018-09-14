package g

import (
	"bytes"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"godev/mymodels/windows-agent/common/model"
	"github.com/toolkits/net"
	"github.com/toolkits/slice"
	"fmt"
)

var (
	Root   string
	logger *log.Logger
)

func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}

type Log_m struct {
	Log_file    string
	Lock        sync.Mutex
	Logfile_obj *os.File
}

func (l Log_m) Write(p []byte) (n int, err error) {
	os.Stdout.Write(p)
	//fmt.Printf("%c\n",p)

	n, e := l.Write_file(p)
	return n, e
}

//todo 如果不绑定指针，l变量在方法内 会被copy一份
func (l *Log_m) Write_file(p []byte) (n int, err error) {
	//p=append(p,'\n')
	l.Lock.Lock()
	n, e := l.Logfile_obj.Write(p)
	defer l.Lock.Unlock()
	return n, e
}

func (l *Log_m) Createlogfile() {
	f, err := os.OpenFile(l.Log_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Errorf("err:%s\n", err)
	}
	l.Logfile_obj = f
}

var log1 *log.Logger
var logfile_instance *Log_m

func InitLog() {
	fileName := Config().Logfile
	//logFile, err := os.Create(fileName)
	//if err != nil {
	//	log.Fatalln("open file error !")
	//}

	logfile_instance = &Log_m{Log_file: fileName}
	logfile_instance.Createlogfile()
	logger = log.New(*logfile_instance, "", log.Ldate|log.Ltime|log.Lshortfile)


	//logger = log.New(logFile, "[Debug]", log.LstdFlags)
	log.Println("logging on", fileName)
}

func Logger() *log.Logger {
	lock.RLock()
	defer lock.RUnlock()
	return logger
}

var LocalIps []string

func InitLocalIps() {
	var err error
	LocalIps, err = net.IntranetIP()
	if err != nil {
		logger.Fatalln("get intranet ip fail:", err)
	}
}

var (
	HbsClient *SingleConnRpcClient
)

func InitRpcClients() {
	if Config().Heartbeat.Enabled {
		HbsClient = &SingleConnRpcClient{
			RpcServer: Config().Heartbeat.Addr,
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
	}

	var resp model.TransferResponse
	SendMetrics(metrics, &resp)

	if debug {
		logger.Println("<=", &resp)
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
