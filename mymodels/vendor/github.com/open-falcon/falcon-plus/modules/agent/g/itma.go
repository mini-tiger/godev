package g

import (
	"log"

	"fmt"
	"net"
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"container/list"

	"github.com/open-falcon/falcon-plus/common/model"
	"io/ioutil"
	"bytes"
	"github.com/toolkits/file"
)


func mySplit(s string,ds string) []string {
	var rs []string
	for _,d := range strings.Split(s, ds) {
		if len(d) > 0 {
			rs = append(rs, d)
		}
	}
	return rs;
}

type ProcessAppsysWorker interface {
    GetAppsys() string
    GetType() string
    Find(line string) bool
}


type EnvGridConfig struct {
	Debug     		bool
	JsonConfig		model.EnvGridConfigResponse
	ConfigInterval  int
	DataInterval	int
	Appsys      	[]ProcessAppsysWorker
}



var (
	envGridConfig   EnvGridConfig
	hardware 		bool
	manufacturer 	string
	productName 	string
	version 		string
	serialNumber 	string
)


func GetEnvGridConfig() *EnvGridConfig {
	return &envGridConfig
}





type AppsysProcess struct {
    type_ string
    pids *list.List
    pnls *list.List
}    
func (ap AppsysProcess) addProcess(pid string) {
	ap.pids.PushBack(pid)
}
func (ap AppsysProcess) addProcessNetLink(pnl ProcessNetLink) {
	ap.pnls.PushBack(pnl)
}


type ProcessNetLink struct {
    localPort string
    remoteIp string
    remotePort string
}



type TraceRouteLink struct {
	remote string
	trace *list.List
}
func (trl TraceRouteLink) addTrace(ip string) {
	trl.trace.PushBack(ip)
}
func (trl TraceRouteLink) fomatLast() {
	if trl.trace.Len() == 0 { return }
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
func GetHardware() (string, string, string, string) {
	if hardware {
		return manufacturer, productName, version, serialNumber
	}

	rs, err := readLine("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		fmt.Println(err)
	} else {
		manufacturer = rs
	}
	rs, err = readLine("/sys/class/dmi/id/product_name")
	if err != nil {
		fmt.Println(err)
	} else {
		productName = rs
	}
	rs, err = readLine("/sys/class/dmi/id/product_version")
	if err != nil {
		fmt.Println(err)
	} else {
		version = rs
	}
	rs, err = readLine("/sys/class/dmi/id/product_serial")
	if err != nil {
		fmt.Println(err)
	} else {
		serialNumber = rs
	}

	hardware = true

	return manufacturer, productName, version, serialNumber
}
func GetHardwareBydmidecode() (string, string, string, string)  {
	if hardware {
		return manufacturer, productName, version, serialNumber
	}

	cmd := exec.Command("/bin/bash", "-c", "dmidecode -t system")
    stdout, err := cmd.StdoutPipe()
    if err != nil {
		fmt.Println(err)
	}
    cmd.Start()
    reader := bufio.NewReader(stdout)
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
        	fmt.Println("E:", err2)
            break
        }
        
        if manufacturer == "" {
        	idx := strings.Index(line, "Manufacturer: ")
        	if idx > 0 {
        		manufacturer = line[idx + 14:len(line)-1]
        	}
        }
        if productName == "" {
        	idx := strings.Index(line, "Product Name: ")
        	if idx > 0 {
        		productName = line[idx + 14:len(line)-1]
        	}
        }
        if version == "" {
        	idx := strings.Index(line, "Version: ")
        	if idx > 0 {
        		version = line[idx + 9:len(line)-1]
        	}
        }
        if serialNumber == "" {
        	idx := strings.Index(line, "Serial Number: ")
        	if idx > 0 {
        		serialNumber = line[idx + 15:len(line)-1]
        	}
        }
    }
 	cmd.Wait()

 	hardware = true

 	return manufacturer, productName, version, serialNumber
}
func EnvGrid() model.EnvGrid {
	var process2appsys []ProcessAppsysWorker
	/*/	
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"MySQL", "4", "and", []string{"mysqld"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Apache", "0", "and", []string{"httpd"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Memcached", "0", "and", []string{"memcached"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Tomcat", "0", "and", []string{"java", "tomcat"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Java", "0", "or", []string{"java", "javaw"}}))
	process2appsys = append(process2appsys, ProcessAppsysWorker(DefaultProcessAppsysWorker{"Ssh", "0", "and", []string{"ssh"}}))
	/*/
	process2appsys = GetEnvGridConfig().Appsys
   	log.Printf("%s",process2appsys,)

	var appsys2pid map[string]AppsysProcess
	appsys2pid = make(map[string]AppsysProcess)
	
	var remotes map[string]TraceRouteLink
	remotes = make(map[string]TraceRouteLink)

	ipaddress := []string {}
	addrs, err := net.InterfaceAddrs();
	if err != nil {
		fmt.Println(err)
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
	hostname,_ := os.Hostname()
    manufacturer, productName, version, serialNumber := GetHardware()

    cmd := exec.Command("/bin/bash", "-c", "ps -eo pid,user,state,tty,command")
    stdout, err := cmd.StdoutPipe()
    cmd.Start()
    reader := bufio.NewReader(stdout)
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
        	fmt.Println("E:", err2)
            break
        }
        for _,worker := range process2appsys {
            if worker.Find(line) {
            	ap,ok := appsys2pid[worker.GetAppsys()]
            	if !ok {
            		ap = AppsysProcess{worker.GetType(), list.New(), list.New()}
            		appsys2pid[worker.GetAppsys()] = ap
            	}
            	ap.addProcess(mySplit(line, " ")[0])
            }
        }
    }
 	cmd.Wait()
    


    cmd = exec.Command("/bin/bash", "-c", "netstat -tunap | grep -v 0.0.0.0:* | grep -v :::* | grep -v 127.0.0.1:")
    stdout, err = cmd.StdoutPipe()
    cmd.Start()
    reader = bufio.NewReader(stdout)
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
            break
        }
        for _,v := range appsys2pid {
        	for e := v.pids.Front(); e != nil; e = e.Next() {
				pid := e.Value
        		if !strings.Contains(line,fmt.Sprintf(" %s/", pid)) { continue }
        		cfg := mySplit(line," ")
    			if len(cfg) < 4 { continue }
			    if !strings.Contains(cfg[3],":") { continue }
			    if !strings.Contains(cfg[4],":") { continue }
			    
			    pnl := ProcessNetLink{strings.Split(cfg[3],":")[1], strings.Split(cfg[4],":")[0], strings.Split(cfg[4],":")[1]}
			    v.addProcessNetLink(pnl)
			    r := strings.Split(cfg[4],":")[0]
			    _,ok := remotes[r]
			    if !ok { remotes[r] = TraceRouteLink{r, list.New()} }
	        }
        }
    }
 	cmd.Wait()


	for remote,trl := range remotes {
	    cmd = exec.Command("/bin/bash", "-c", "traceroute -n -w 5 -m 15 -q 1 "+remote)
	    stdout, err = cmd.StdoutPipe()
	    cmd.Start()
	    reader = bufio.NewReader(stdout)
	    for {
	        line, err2 := reader.ReadString('\n')
	        if err2 != nil || io.EOF == err2 {
	            break
	        }
    
    		if strings.Contains(line,"traceroute") { continue }
    		data := mySplit(line," ")
    		if len(data) < 3 { continue }
    		if data[1] == "*" { continue }
    		trl.addTrace(data[1])
    	}
 		cmd.Wait()
    		
  		trl.fomatLast();
  	}
    
    
    
    var appsyss []model.Appsys
	for k,v := range appsys2pid {
		var links []model.Link
    	for e := v.pnls.Front(); e != nil; e = e.Next() {
    		var pnl ProcessNetLink = e.Value.(ProcessNetLink)
    		links = append(links,model.Link{pnl.localPort,pnl.remoteIp,pnl.remotePort})
		}
		appsyss = append(appsyss,model.Appsys{k,v.type_,links})
	}
	var routes []model.Route
	for k,v := range remotes {
		var traces []string
    	for e := v.trace.Front(); e != nil; e = e.Next() {
    		traces = append(traces,e.Value.(string))
      	}
      	routes = append(routes,model.Route {k,traces})
    }
    
    env := model.EnvGrid{
		Hostname:   hostname,
		Address:    ipaddress,
		Manufacturer:   manufacturer,
		ProductName:    productName,
		Version:    version,
		SerialNumber:    serialNumber,
		Appsyss:  	appsyss,
		Routes: 	routes,
	}
    
    return env
}