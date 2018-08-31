package g

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"fmt"
	"os/exec"
	"log"
	"bufio"
	"io"
	"strings"
	"strconv"
)

//config
type EnvPortProcessConfig struct {
	Debug          bool
	JsonConfig     model.PortprocessConfResponse
	ConfigInterval int
	DataInterval   int
}

var EnvPortConfig EnvPortProcessConfig

func NewPortProcessConfig() *EnvPortProcessConfig {
	return &EnvPortConfig
}

//
//type Portprocess_task_env struct {
//	PP []model.PortProcessEnv
//}

//var Portprocess_result_env Portprocess_result

//func NewPortprocess_result() *Portprocess_result{
//	return &Portprocess_result_env
//}

//result map[int]struct{process:nil,type:nil} map[port]

//var Portprocess_search_env Portprocess_task_env

func Getportprocess_data() (model.Portprocess_result) {
	gp := EnvPortConfig

	portprocess_result_env := model.Portprocess_result{}

	for _, v := range gp.JsonConfig.Portprocess_slice {
		ips, err := ips_business(v)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(ips)

		pidsi, err := pids_business(v)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(pidsi)

		temp := model.Portprocess_sub_result{
			Type:     v.Type,
			Port:     v.Port,
			Pids:     pidsi,
			Ip_route: ips,
		}
		//fmt.Println("=======================================")
		//for k,v:=range temp.Ip_route {
		//	if v.Back() != nil {
		//		fmt.Println(k, v.Back().Value)
		//	}
		//}
		portprocess_result_env.Pr = append(portprocess_result_env.Pr, &temp)

	}

	return portprocess_result_env
}

func pids_business(v model.PortProcessEnv) ([]int, error) {
	cmd_string := fmt.Sprintf("lsof -i tcp:%d|awk '{print $2}'|uniq|grep -v 'PID'", v.Port)
	fmt.Println("bash", "-c", cmd_string)
	cmd := exec.Command("/bin/bash", "-c", cmd_string)

	pids := make([]string, 0)

	var pidsi []int = make([]int, 0)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("ips_business run port:%d, err:%v", v.Port, err)
		return pidsi, err
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(strings.Trim(line, "\n"))

		//fmt.Println([]byte(line))
		pids = append(pids, line)
	}

	cmd.Wait()

	pidsi = func(s []string) ([]int) { //conv  string -> int
		for _, v := range s {
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}

			pidsi = append(pidsi, i)
			//fmt.Printf("%v,%d\n",v,s)
		}
		return pidsi

	}(pids)

	return pidsi, nil
}

func ips_business(v model.PortProcessEnv) (map[string]model.Route_link, error) {
	cmd_string := fmt.Sprintf("lsof -i tcp:%d|grep ESTABLISHED|awk '{print $9}'|awk -F '->' '{print $2}'|grep -v local|cut -d : -f 1|uniq", v.Port)
	fmt.Println("bash", "-c", cmd_string)
	cmd := exec.Command("/bin/bash", "-c", cmd_string)
	//ips := make([]string, 1)
	ips_route := make(map[string]model.Route_link)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("ips_business run port:%d, err:%v", v.Port, err)
		return ips_route, err
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(strings.TrimSpace(line), "\n")
		//fmt.Println([]byte(line))
		//ips = append(ips, line)
		ips_route[line] = route_link(v.Port, line)
	}

	cmd.Wait()

	return ips_route, nil
}

func route_link(port int, ip string) (model.Route_link) {
	cmd_string := fmt.Sprintf("traceroute -n -w 5 -m 15 -q 1 -p %d -4 %s|awk '{print $1,$2}'", port, ip)
	fmt.Println("bash", "-c", cmd_string)
	cmd := exec.Command("/bin/bash", "-c", cmd_string)

	//route := model.Route_link{}

	route_links := model.Newroute_list()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("route_link run port:%d, ip:%s,err:%v", port, ip, err)
		return *route_links
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		//fmt.Printf(line)
		//time.Sleep(2*time.Second)
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, "traceroute to") || strings.Contains(line, "*") {
			continue
		}
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "\n")
		//route_links.Lock.Lock()
		//fmt.Println("==================================",line)
		//route_links.PushBack(strings.TrimSpace(line))
		//route_links.Links = append(route_links.Links,line)
		//route_links.Lock.Unlock()
		split_data := strings.Split(line, " ")
		if len(split_data) < 2 {
			continue
		}
		fmt.Printf("%T,%s\n", split_data[0],split_data[0])
		key:=strings.TrimSpace(split_data[0])
		k,err:=strconv.Atoi(key)
		if err !=nil{
			continue
		}
		route_links.L[k] = split_data[1]
	}
	cmd.Wait()
	//if route_links.Links.Back() != nil{
	//	fmt.Println("--------------------------------")
	//	fmt.Println(route_links.Links.Back().Value)
	//	fmt.Println(route_links.Links.Back().Value)
	//	fmt.Println(route_links.Links.Back().Value)
	//	fmt.Println(route_links.Links.Len())
	//}

	return *route_links

}
