package g

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"tjtools/file"
	"godev/mymodels/windows-agent/common/model"
	"io"
	"io/ioutil"
	rnet "net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"runtime"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"strconv"
)

func mySplit(s string, ds string) []string {
	var rs []string
	for _, d := range strings.Split(s, ds) {
		if len(d) > 0 {
			rs = append(rs, d)
		}
	}
	return rs
}

type ProcessAppsysWorker interface {
	GetAppsys() string
	GetType() string
	Find(line string) bool
}

type EnvGridConfig struct {
	Debug          bool
	JsonConfig     model.EnvGridConfigResponse
	ConfigInterval int
	DataInterval   int
	Appsys         []ProcessAppsysWorker
}

var (
	envGridConfig EnvGridConfig
	hardware      bool
	manufacturer  string
	productName   string
	version       string
	serialNumber  string
)

func GetEnvGridConfig() *EnvGridConfig {
	return &envGridConfig
}

type AppsysProcess struct {
	type_ string
	pids  *list.List
	pnls  *list.List
}

func (ap AppsysProcess) addProcess(pid string) {
	ap.pids.PushBack(pid)
}
func (ap AppsysProcess) addProcessNetLink(pnl ProcessNetLink) {
	ap.pnls.PushBack(pnl)
}

type ProcessNetLink struct {
	localPort  string
	remoteIp   string
	remotePort string
}

type TraceRouteLink struct {
	remote string
	trace  *list.List
}

func (trl TraceRouteLink) addTrace(ip string) {
	trl.trace.PushBack(ip)
}
func (trl TraceRouteLink) fomatLast() {
	if trl.trace.Len() == 0 {
		return
	}
	if trl.trace.Back().Value == trl.remote {
		trl.trace.Remove(trl.trace.Back())
	}
}

func readLine(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(bytes.NewBuffer(bs))
	line, err := file.ReadLine(reader)
	if err != nil {
		return "", err
	}
	return string(line), nil
}

func gethare_sub(sub_info string) string {
	temps := ""
	cmd_string := fmt.Sprintf("wmic BaseBoard get %s", sub_info)
	cmd := exec.Command("cmd", "/c", cmd_string)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Println(err)
		logger.Printf("Error get hardinfo:%s,%s\n", sub_info, err)
		return temps
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.Contains(strings.ToLower(string(line)), sub_info) {
			continue
		}
		temps = strings.TrimSpace(string(line))
		temps = strings.Trim(temps, "\n")
		break
	}

	return temps
}

func GetHardware() (string, string, string, string) {
	if hardware {
		return manufacturer, productName, version, serialNumber
	}
	manufacturer = gethare_sub("manufacturer")
	productName = gethare_sub("product")
	version = gethare_sub("version")
	serialNumber = gethare_sub("serialnumber")

	hardware = true

	return manufacturer, productName, version, serialNumber
}

func EnvGrid() model.EnvGrid {
	//var process2appsys []ProcessAppsysWorker
	/*/
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"MySQL", "4", "and", []string{"mysqld"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Apache", "0", "and", []string{"httpd"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Memcached", "0", "and", []string{"memcached"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Tomcat", "0", "and", []string{"java", "tomcat"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Java", "0", "or", []string{"java", "javaw"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Ssh", "0", "and", []string{"ssh"}}))
	/*/
	//process2appsys = GetEnvGridConfig().Appsys
	//logger.Printf("%s",process2appsys,)

	//var appsys2pid map[string]AppsysProcess
	//appsys2pid = make(map[string]AppsysProcess)
	//
	//var remotes map[string]TraceRouteLink
	//remotes = make(map[string]TraceRouteLink)

	err, ipaddress := GetIps()
	if err != nil {
		logger.Printf("GetIps FAil err:%s\n", err)
	}
	//ipaddress := []string{}
	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//
	//	logger.Printf("net.interface err: %s\n", err)
	//} else {
	//	for _, address := range addrs {
	//
	//		// 检查ip地址判断是否回环地址
	//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//			if ipnet.IP.To4() != nil {
	//				ipaddress = append(ipaddress, ipnet.IP.String())
	//			}
	//
	//		}
	//	}
	//}
	hostname, _ := os.Hostname()
	manufacturer, productName, _, serialNumber := GetHardware()
	//fmt.Println(manufacturer, productName, version, serialNumber)

	var appsyss []model.Appsys
	//for k,v := range appsys2pid {
	//var links []model.Link
	//	for e := v.pnls.Front(); e != nil; e = e.Next() {
	//		var pnl ProcessNetLink = e.Value.(ProcessNetLink)
	//		links = append(links,model.Link{pnl.localPort,pnl.remoteIp,pnl.remotePort})
	//}
	//appsyss = append(appsyss,model.Appsys{k,v.type_,links})
	//}
	var routes []model.Route
	//for k,v := range remotes {
	//var traces []string
	//	for e := v.trace.Front(); e != nil; e = e.Next() {
	//		traces = append(traces,e.Value.(string))
	//  	}
	//  	routes = append(routes,model.Route {k,traces})
	//}

	//extinfo := getExtInfo()
	env := model.EnvGrid{
		Hostname: hostname,
		//Address:    ipaddress,
		Manufacturer: manufacturer,
		ProductName:  productName,
		//Version:      version,
		Version:      getOSVersion(),
		SerialNumber: serialNumber,
		Appsyss:      appsyss,
		Routes:       routes,
		Extinfo:      getExtInfo(ipaddress[0]),
		//AgentInfo:    GetAgentVer(),
		AgentInfo: model.AgentInfo{"v1"},
	}
	//fmt.Println(env.Extinfo)
	return env
}

func getExtInfo(ipaddress string) (tmp model.EnvGridExt) {
	//var tmp *model.EnvGridExt
	//tmp.Osname = runtime.GOOS
	tmp.Osname = *getOSName()
	tmp.Osbit = strings.Split(string(runtime.GOARCH), "amd")[1]
	tmp.MAC = rMac(OutIP)
	tmp.Disktotal = getDiskTotal()
	tmp.Memtotal = getMemTotal()
	tmp.Cpucores = runtime.NumCPU()
	tmp.Cpumhz, tmp.Cpumodel = getCpuInfo()
	tmp.UUID = Uuid
	tmp.Biz = BizId
	//tmp.Cpumodel = getCpuModel()
	tmp.Address = GetIP(ipaddress) //  todo 以通过TCP链接获取到的IP为准
	return tmp
}

func GetIP(ip string) (string) {
	if *OutIP != "" {
		return *OutIP
	} else {
		return ip
	}
}

func TrimAny(cm *string) *string {
	*cm = strings.Trim(*cm, " ")
	*cm = strings.Trim(*cm, "\t")
	*cm = strings.Trim(*cm, "\r\n")
	return cm
}

func getOSVersion() (string) {
	n, err := host.Info()
	if err != nil {
		logger.Error("获取主机操作系统版本失败:%s", err)
	}

	return n.PlatformVersion
}

func getOSName() (cm *string) {
	n, err := host.Info()
	if err != nil {
		logger.Error("获取主机操作系统名称失败:%s", err)
	}

	return &n.Platform
}

//
//func getCpuModel() (cm string) {
//	cmd := exec.Command("/bin/bash", "-c", "grep -i 'model name' /proc/cpuinfo |uniq|cut -d ':' -f 2")
//	buf, _ := cmd.CombinedOutput()
//	cmd.Run()
//	cm = strings.Trim(string(buf), " ")
//	cm = strings.Trim(cm, "\r\n")
//	return
//}

func getCpuInfo() (cpuMhz int, cpuModel string) {
	if m,err:=getCpusub("maxclockspeed"); err ==nil{
		cpuMhz,err=strconv.Atoi(m)
		if err!=nil{
			logger.Error("cpuMhz 转换为数字失败 %s",err)
		}
	}
	cpuModel,err:=getCpusub("name")
	if err !=nil{
		logger.Error("cpuModel 获取失败 %s",err)
	}
	return
}


func getCpusub(sub_info string) (temps string,err error) {
	//temps := ""
	cmd_string := fmt.Sprintf("wmic cpu get %s\n",sub_info)
	cmd := exec.Command("cmd", "/c", cmd_string)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//logger.Printf("Error get hardinfo:%s,%s\n", sub_info, err)
		return "",err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.Contains(strings.ToLower(string(line)), sub_info) {
			continue
		}
		temps = strings.TrimSpace(string(line))
		temps = strings.Trim(temps, "\n")
		break
	}

	return temps,nil
}

func getMemTotal() int64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		logge.Error("get memTotal err :%s", err)
	}
	return int64(v.Total) / 1024 / 1024 / 1024
}

func getDiskTotal() (int64) {
	d, err := disk.Usage("/")
	if err != nil {
		logger.Error("get disk err:%s", err)
	}
	return int64(d.Total) / 1024 / 1024 / 1024
}

func rMac(ip *string) (string) {
	ni,_:=net.Interfaces()
	for _,v:=range ni{
		tl := len(v.Addrs)
		switch {
		case tl==0:
			continue
		case tl==1:
			if strings.Contains(v.Addrs[0].Addr,*ip) {
				return v.HardwareAddr
			}
		case tl>=2:
			for _,vv:=range v.Addrs{
				if strings.Contains(vv.Addr,*ip) {
					return v.HardwareAddr
				}
			}
		}


		//fmt.Printf("%+v\n",v)

	}
	return  ""
}

func GetIps() (err error, ipaddress []string) {
	//ipaddress := []string {}
	addrs, err := rnet.InterfaceAddrs()
	if err != nil {
		logge.Error("%s", err)
		return
	} else {
		for _, address := range addrs {

			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*rnet.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipaddress = append(ipaddress, ipnet.IP.String())
				}

			}
		}
	}
	return
}
