package g

import (
	"bufio"
	"container/list"
	"fmt"
	"net"
	"os"
	"strings"

	"bytes"
	"github.com/toolkits/file"
	"godev/mymodels/windows-agent/common/model"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
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

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		log.Printf("Error get hardinfo:%s,%s\n", sub_info, err)
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
	//log.Printf("%s",process2appsys,)

	//var appsys2pid map[string]AppsysProcess
	//appsys2pid = make(map[string]AppsysProcess)
	//
	//var remotes map[string]TraceRouteLink
	//remotes = make(map[string]TraceRouteLink)

	ipaddress := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		log.Printf("net.interface err: %s\n", err)
	} else {
		for _, address := range addrs {

			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipaddress = append(ipaddress, ipnet.IP.String())
				}

			}
		}
	}
	hostname, _ := os.Hostname()
	manufacturer, productName, version, serialNumber := GetHardware()
	fmt.Println(manufacturer, productName, version, serialNumber)

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

	env := model.EnvGrid{
		Hostname:     hostname,
		Address:      ipaddress,
		Manufacturer: manufacturer,
		ProductName:  productName,
		Version:      version,
		SerialNumber: serialNumber,
		Appsyss:      appsyss,
		Routes:       routes,
	}

	return env
}
