package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	conn, err := net.Dial("tcp", "192.168.1.108:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//go mustCopy(os.Stdout, conn)
	j := jsonStr{"a", 1}
	bytejson, err := json.Marshal(j)
	if err != nil {
		log.Printf("json marshal err:%s\n", err)
	}

	handleWrite(conn, bytejson)
	//b := bytes.NewBuffer(bytejson)
	mustCopy(conn, os.Stdin)
}

func handleWrite(conn net.Conn, bj []byte) {

	//_, e := conn.Write([]byte("hello1 " + "\r\n")) // 如果写入多条，需要服务端 多次read
	_, e := conn.Write(bj)
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())

	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
