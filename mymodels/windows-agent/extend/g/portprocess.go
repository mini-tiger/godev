package g

import (
	"bufio"
	"fmt"
	"godev/mymodels/windows-agent/common/model"
	extned_funcs "godev/mymodels/windows-agent/extend/funcs"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
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

var Cmdline []string

func create_cmdline_data() error {
	cmd_string := fmt.Sprintf("netstat -ano")
	//fmt.Println("cmd", "/c", cmd_string)
	cmd := exec.Command("cmd", "/c", cmd_string)

	stdout, err := cmd.StdoutPipe()
	cmd.Start()

	if err != nil {
		log.Fatalf("err:%v", err)
		return err
	}
	reader := bufio.NewReader(stdout)
	Cmdline = make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(strings.TrimSpace(line), "\n")
		Cmdline = append(Cmdline, line)
	}

	cmd.Wait()

	return nil
}

func Getportprocess_data() model.Portprocess_result {
	gp := EnvPortConfig

	portprocess_result_env := model.Portprocess_result{}

	for _, v := range gp.JsonConfig.Portprocess_slice {

		err := create_cmdline_data()  //除了路由以外，都使用netstat -ano为数据基础
		//fmt.Println(Cmdline)
		//fmt.Println("==================111111")
		if err != nil {

			return portprocess_result_env
		}
		//fmt.Println("==================222222")
		pidsi, err := pids_business(v)
		if err != nil {
			log.Printf("get port: %d, pid err: %s", v.Port, err)
		}
		//fmt.Println(pidsi)
		//
		ips, err := ips_business(v)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(ips)

		temp := model.Portprocess_sub_result{
			Type:     v.Type,
			Port:     v.Port,
			Pids:     pidsi,
			Ip_route: ips,
		}

		portprocess_result_env.Pr = append(portprocess_result_env.Pr, &temp)

	}

	return portprocess_result_env
}

func pids_business(v model.PortProcessEnv) ([]int, error) {
	log.Printf("begin check pids port : %d \n", v.Port)

	pids := make([]string, 0)

	var pidsi []int = make([]int, 0)

	for _, line := range Cmdline {

		//line = strings.TrimSpace(strings.Trim(line, "\n"))
		tmp_data_slice := strings.Fields(line)
		if len(tmp_data_slice) < 5 {
			continue
		}
		if tmp_data_slice[3] != "ESTABLISHED" {
			continue
		}

		ip_port := strings.Split(tmp_data_slice[1], ":")
		if ip_port[1] == fmt.Sprintf("%d", v.Port) {
			//fmt.Println(ip_port)
			//fmt.Printf("pid: %s\n",tmp_data_slice[len(tmp_data_slice)-1])
			pids = append(pids, tmp_data_slice[len(tmp_data_slice)-1])
		}

	}

	pidsi = func(s []string) []int { //conv  string -> int
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
	log.Printf("end check pids port : %d \n", v.Port)

	return extned_funcs.Set(pidsi), nil
}

func ips_business(v model.PortProcessEnv) (map[string]model.Route_link, error) {
	log.Printf("begin check ips port : %d \n", v.Port)
	//cmd_string := fmt.Sprintf("lsof -i tcp:%d|grep ESTABLISHED|awk '{print $9}'|awk -F '->' '{print $2}'|grep -v local|cut -d : -f 1|uniq", v.Port)
	//fmt.Println("bash", "-c", cmd_string)
	//cmd := exec.Command("/bin/bash", "-c", cmd_string)
	//ips := make([]string, 1)
	ips_route := make(map[string]model.Route_link)
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Fatalf("ips_business run port:%d, err:%v", v.Port, err)
	//	return ips_route, err
	//}
	//cmd.Start()

	//reader := bufio.NewReader(stdout)

	for _, line := range Cmdline {

		//line = strings.TrimSpace(strings.Trim(line, "\n"))
		tmp_data_slice := strings.Fields(line)
		if len(tmp_data_slice) < 5 {
			continue
		}
		if tmp_data_slice[3] != "ESTABLISHED" {
			continue
		}
		ip_port := strings.Split(tmp_data_slice[1], ":")
		if ip_port[1] == fmt.Sprintf("%d", v.Port) { //port 匹配

			tmp := strings.Split(tmp_data_slice[2], ":")
			if tmp[0] == "127.0.0.1" {
				continue
			}
			//fmt.Println(tmp[0])

			ips_route[line] = route_link(v.Port, tmp[0])
		}
	}

	//cmd.Wait()
	log.Printf("end check ips port : %d \n", v.Port)
	return ips_route, nil
}

func route_link(port int, ip string) model.Route_link {
	cmd_string := fmt.Sprintf("tracert -d -h 15 -w 15 -4 %s", ip)
	fmt.Println("cmd", "/c", cmd_string)

	cmd := exec.Command("cmd", "/c", cmd_string)
	route_links := model.Newroute_list()

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return *route_links
	}
	ss := extned_funcs.ConvertByte2String(stdout, extned_funcs.GB18030)
	//fmt.Println(ss)
	//fmt.Println("===================================")
	ss = strings.TrimSpace(ss)
	ss = strings.Trim(ss, "\n")

	s_sub := "路由"
	if strings.Index(ss, "路由:") != -1 {
		s_sub = fmt.Sprintf("%s:", s_sub)
	}
	ss1 := strings.Split(ss, s_sub)

	ss2 := strings.Split(ss1[1], "跟踪完成。")
	ss = strings.Trim(ss2[0], "\n")
	ss = strings.Trim(ss2[0], "\r")
	ss = strings.TrimSpace(ss2[0])

	ss3 := strings.Split(ss, "\r\n")
	//fmt.Println(len(ss3))
	for _, v := range ss3 {
		sl := strings.Fields(v)

		if strings.Contains(sl[len(sl)-1], "请求超时") {
			continue
		}
		k, err := strconv.Atoi(sl[0])
		if err != nil {
			continue
		}
		route_links.L[k] = sl[len(sl)-1]

	}
	fmt.Println(route_links)
	return *route_links

}
