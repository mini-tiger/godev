package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net"
	"time"
	"strings"
	"log"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"bufio"
	"bytes"
	"github.com/toolkits/file"
)

func GetOutboundIP() *string {
	//L.Add(1)
	//defer func() {
	//	L.Done()
	//}()
	//conn, err := net.Dial("udp", "8.8.8.8:80")
	conn, err := net.DialTimeout("tcp", "192.168.43.11:22", time.Duration(15)*time.Second)
	if err != nil {

		log.Println(err)
	}
	defer conn.Close()

	//localAddr := conn.LocalAddr().(*net.UDPAddr)
	localAddr := conn.LocalAddr().(*net.TCPAddr)
	//fmt.Println(localAddr.IP)
	//return localAddr.IP
	//fmt.Printf("%T,%s\n",localAddr.IP,localAddr.IP)
	tmp := fmt.Sprintf("%s", localAddr.IP)
	tmp = strings.TrimSpace(tmp)
	tmp = strings.Trim(tmp, "\n")
	return &tmp
}

func Md5(raw string) string {
	h := md5.Sum([]byte(raw))
	return hex.EncodeToString(h[:])
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

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s, %T\n", u2, u2)

	// Parsing UUID from string input
	ip := GetOutboundIP() //  获取IP

	sn, err := readLine("/sys/class/dmi/id/product_serial") // 获取sn号
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*ip)
	fmt.Println(sn)

	//m := Md5(fmt.Sprintf("%s%s", sn, *ip)) // todo ip,sn生成md5 字符串 唯一

	mac := getMAC(*ip)
	if mac == "" {
		log.Println("mac err:")
	}else {
		fmt.Println("mac:",mac)
	}


	m := Md5(fmt.Sprintf("%s%s", mac, *ip)) // todo mac,ip生成md5 字符串 唯一

	fmt.Println(m, len(m))
	u3, err := uuid.FromString(m) // 生成UUID格式
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
	}
	fmt.Printf("Successfully parsed: %s\n", u3)
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
				i := strings.Split(ip.String(),"/")
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
