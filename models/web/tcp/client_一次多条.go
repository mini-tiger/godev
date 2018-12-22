package main

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"
)

type jsonStr struct {
	A string
	B int
}

// 获取服务端发送的消息
func clientRead(conn net.Conn) int {
	buf := make([]byte, 5)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("receive server info faild: %s\n", err)
	}
	// string conver int
	off, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		log.Fatalf("string conver int faild: %s\n", err)
	}
	return off
}

// 发送消息到服务端
func clientWrite(conn net.Conn, data []byte) {
	_, err := conn.Write(data)
	if err != nil {
		log.Printf("[ERROR] send 【%s】 content faild: %s\n", string(data), err)
	}
	log.Printf("send 【%s】 content success\n", string(data))
}

// client conn
func clientConn(conn net.Conn) {
	defer conn.Close()
	j := jsonStr{"a", 1}
	bytejson, _ := json.Marshal(j)
	for {

		//data := make([]byte, 10)

		time.Sleep(time.Second * 1)
		clientWrite(conn, bytejson)

	}
}

func main() {
	// connect timeout 10s
	conn, err := net.DialTimeout("tcp", "192.168.1.108:8000", time.Second*10)
	if err != nil {
		log.Fatalf("client dial faild: %s\n", err)
	}
	clientConn(conn)
}
