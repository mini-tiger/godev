package g

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net"
	"strings"
	"os"
	"path/filepath"
	"github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/host"
	"tjtools/file"
)

type Biz struct {
	Bizid int `json:"bizid"`
}

var BizId int
var Uuid string

func Md5(raw string) string {
	h := md5.Sum([]byte(raw))
	return hex.EncodeToString(h[:])
}
func getMAC(ips string) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {

		interaddr, _ := inter.Addrs()
		for _, a := range interaddr {
			//fmt.Println(a)
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				ip := ipnet.IP.To4()
				//fmt.Println(ip.String())
				i := strings.Split(ip.String(), "/")
				//fmt.Println(i)
				if i[0] == ips {
					//fmt.Println(ipnet.IP.String())
					//fmt.Println(inter.Name, inter.HardwareAddr.String())
					return inter.HardwareAddr.String()
				}
			}
		}
	}
	return ""
}

func getBiz() (bizid int) {

	dir, _ := os.Getwd()
	f := filepath.Join(dir, "biz.json")
	configContent, err := file.ToTrimString(f)

	if err != nil {
		logger.Error("转换biz json文件err:%s\n", err)
		//log.Fatalln("read config file:", f, "fail:", err)
		//WLog(fmt.Sprintf("read config file:%s fail:%s", cfg, err))
		return
	}

	var c Biz
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		//log.Fatalln("parse config file:", f, "fail:", err)
		//WLog(fmt.Sprintf("parse config file: %s fail: %s", cfg, err))
		logger.Error("解析biz json文件err:%s\n", err)
		return
	}
	bizid = c.Bizid
	return
}
func LoadUUIDBIZ()  {

	bizid := getBiz() // 这里需要扫描文件
	mac := getMAC(*OutIP)
	if mac == "" {
		logger.Error("mac err:")
		return
	} else {

		logger.Debug("获取MAC地址:%s", mac)
	}

	//m := Md5(fmt.Sprintf("%s%s", mac, *OutIP)) // mac,ip 生成md5 字符串 唯
	m,_:=host.Info()
	u, err := uuid.FromString(m.HostID)               // 生成UUID格式
	if err != nil {
		logger.Error("生成uuid失败,ip:%s,MAC:%s,err:%s\n", *OutIP, mac, err)
		return
	}
	logger.Printf("根据IP,MAC第一次生成uuid成功,uuid:%s\n", u.String())

	BizId = bizid
	logger.Printf("BizId:%d\n",BizId)
	Uuid = u.String()
}