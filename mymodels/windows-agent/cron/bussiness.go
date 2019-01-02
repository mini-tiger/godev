package cron

import (
	"fmt"
	"godev/mymodels/windows-agent/g"
	"time"
	"net"
	"strings"
	"regexp"
)

func CollectInfo() {
	GetOutboundIP()
	//go GetOutPubIP()
}

//func GetOutPubIP() { // 连接公网的出网地址
//	for {
//		var tmp string
//		par := `(\d+)\.(\d+)\.(\d+)\.(\d+)`
//		func() {
//			rq := httplib.Get("http://myexternalip.com/raw").SetTimeout(time.Duration(5)*time.Second, time.Duration(5)*time.Second)
//			resp, err := rq.Response()
//			if err != nil {
//				g.Logger().Error("获取本机出网公网IP:", err)
//				return
//			}
//			//defer conn.Close()
//			content, _ := ioutil.ReadAll(resp.Body)
//			tmp = string(content)
//		}()
//
//		tmp = strings.TrimSpace(tmp)
//		tmp = strings.Trim(tmp, "\n")
//		g.OutPubIP = tmp
//		if ok, _ := regexp.MatchString(par, g.OutPubIP); ok {
//			g.Logger().Printf("获取本机出网公网IP: %s,成功，退出循环\n", g.OutPubIP)
//			break
//		} else {
//			time.Sleep(time.Duration(10) * time.Second)
//		}
//	}
//
//}

func GetOutboundIP() { // 连接公网的内网地址
	var tmp string
	par := `(\d+)\.(\d+)\.(\d+)\.(\d+)`
	func() {
		conn, err := net.DialTimeout("tcp", g.Config().Ubs.Addr, time.Duration(15)*time.Second)
		if err != nil {
			g.Logger().Error("获取本机出网私网IP err:%s", err)
			return
		}
		//defer conn.Close()

		//localAddr := conn.LocalAddr().(*net.UDPAddr)
		localAddr := conn.LocalAddr().(*net.TCPAddr)
		//fmt.Println(localAddr.IP)
		//return localAddr.IP
		//fmt.Printf("%T,%s\n",localAddr.IP,localAddr.IP)
		tmp = fmt.Sprintf("%s", localAddr.IP)
	}()

	tmp = strings.TrimSpace(tmp)
	tmp = strings.Trim(tmp, "\n")
	g.OutIP = &tmp
	if ok, _ := regexp.MatchString(par, *g.OutIP); ok {
		g.Logger().Printf("获取本机出网私网IP: %s,成功，退出循环\n", *g.OutIP)

	}
}
