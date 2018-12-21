package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"encoding/json"
)
type jsonStr struct {
	A string
	B int
}
func echo(c net.Conn, shout string, delay time.Duration) {

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleRead(conn net.Conn) {
	reader := bufio.NewReader(conn)
	//line, err := reader.ReadString(byte('\n'))
	p:=make([]byte,128)
	len, err:=reader.Read(p)
	log.Printf("recv data:%s,Len:%d\n",string(p),len)

	p=p[0:len]  // todo 接收到的数据，长度不确定，需要截取

	jj := jsonStr{}
	err=json.Unmarshal(p,&jj)
	if err!=nil{
		log.Println("json unmarshal err:",err)
	}
	fmt.Println("json:",jj)

	if err != nil {
		fmt.Printf("Error to read message because of %s\n", err)
		return
	}
	//fmt.Print(line)

}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	handleRead(c)
	fmt.Printf("收到ip:%s 的链接\n", c.RemoteAddr().String())

	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "192.168.1.108:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
