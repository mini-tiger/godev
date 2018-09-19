package main

import (
	"net"
	"time"
	"fmt"
	"strings"
)

func main()  {
	conn, err := net.DialTimeout("tcp", "www.baidu.com:80", 2 * time.Second)
	if err != nil {

	}
	fmt.Println(conn.LocalAddr())
	ip:=conn.LocalAddr().String()
	//fmt.Printf("%T\n",ip)
	ips:=strings.Split(ip,":")
	//fmt.Println(ips[0])
	}