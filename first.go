package main

import (
	"github.com/shirou/gopsutil/cpu"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"time"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"

	"os/exec"
	"syscall"
	"github.com/jander/golog/logger"
	"bufio"
	"io"
	"strings"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")
	fmt.Printf("        Mem       : %v MB  Free: %v MB Used:%v Usage:%f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	if len(c) > 1 {
		for _, sub_cpu := range c {
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			mhz :=sub_cpu.Mhz
			fmt.Printf("        CPU       : %v   %v cores ,mhz:%v \n", modelname, cores,mhz)
		}
	} else {
		sub_cpu := c[0]
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		mhz :=sub_cpu.Mhz
		fmt.Printf("        CPU       : %v   %v cores ,mhz:%v \n", modelname, cores,mhz)
	}
	fmt.Printf("        Network: %v, %v bytes / %v bytes\n",nv[0].Name, nv[0].BytesRecv, nv[0].BytesSent)
	fmt.Printf("        SystemBoot:%v\n", btime)
	fmt.Printf("        CPU Used    : used %f%% \n", cc[0])
	fmt.Printf("        HD        : %v GB  Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("        OS        : %+v\n", n)
	fmt.Printf("        Hostname  : %v  \n", n.Hostname)

	fmt.Println(gethare_sub("name"))
	fmt.Println(gethare_sub("maxclockspeed"))
	var ip string
	ip = "192.168.254.163"
	fmt.Println(rMac(&ip))
}

func rMac(ip *string) (string) {
	fmt.Println(*ip)
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



func gethare_sub(sub_info string) string {
	temps := ""
	cmd_string := fmt.Sprintf("wmic cpu get %s\n",sub_info)
	cmd := exec.Command("cmd", "/c", cmd_string)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
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