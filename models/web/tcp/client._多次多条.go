package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"time"
)

/*
C:\godev>go run client_echo.go
hello
         HELLO
         hello
         hello
*/
type jsonStr struct {
	A string
	B int
}

func main() {
	j := jsonStr{"a", 1}
	bytejson, err := json.Marshal(j)

	if err != nil {
		log.Printf("json marshal err:%s\n", err)
	}
	for {
		SendData("192.168.1.108:8000", bytejson)
		time.Sleep(time.Duration(2) * time.Second)
	}

}

func SendData(addr string, bytejson []byte) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//go mustCopy(os.Stdout, conn)

	handleWrite(conn, bytejson) // todo 一次性发送
}

func handleWrite(conn net.Conn, bj []byte) {

	_, e := conn.Write(bj)
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())

	}

}
