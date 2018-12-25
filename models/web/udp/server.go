package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

//var host = flag.String("host", "", "host")
//var port = flag.String("port", "37", "port")
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data) // todo 读取client 发送过来的
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	daytime := time.Now().Unix()
	fmt.Printf("接收到client addr:%s,数据长度:%d,数据:%s\n", remoteAddr, n, string(data[:n]))
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP(b, remoteAddr) // todo 给client发送
}
