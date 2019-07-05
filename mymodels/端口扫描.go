package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"tjtools/nmap"
	"sync"
)

var ipaddrChan chan string = make(chan string, 0)
var wg sync.WaitGroup = sync.WaitGroup{}
var sm *nmap.SafeMap = nmap.NewSafeMap()

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//os.Exit(1)
	}
}

func tcpConn(addr string) {
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", "192.168.43.14:3389") //获取一个TCP地址信息,TCPAddr
	//checkError(err)
	defer wg.Done()
	conn, err := net.DialTimeout("tcp", addr, time.Duration(2)*time.Second) //创建一个TCP连接:TCPConn
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		sm.Put(addr, fmt.Sprintf("ipadd:%s,ERROR:%s:", addr, err.Error()))
	}else{
		defer conn.Close()
		sm.Put(addr, fmt.Sprintf("success,port:%s\n", addr))

	}

	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) //发送HTTP请求头
	//checkError(err)
	//result, err := ioutil.ReadAll(conn) //获得返回数据
	//checkError(err)
	//fmt.Println(string(result))
}

func checkIP(sm *nmap.SafeMap) {
	for {
		select {
		case addr, ok := <-ipaddrChan:
			if ok {
				//result:=tcpConn(addr)
				//sm.Put(addr,result)
				go tcpConn(addr)
			} else {
				return
			}
		}
	}
}

func main() {
	beginTime := time.Now()
	ip := "192.168.43.14"

	go checkIP(sm)
	for _, port := range []int{21, 80, 135, 3389, 1, 2, 3, 4, 5, 6, 7} {
		//tcpConn(fmt.Sprintf("%s:%d", ip, port))
		//sm.Put(fmt.Sprintf("%s:%d", ip, port), "")
		wg.Add(1)
		ipaddrChan <- fmt.Sprintf("%s:%d", ip, port)
	}
	close(ipaddrChan)
	wg.Wait()
	fmt.Println(sm)
	fmt.Println(time.Now().Sub(beginTime))
}
