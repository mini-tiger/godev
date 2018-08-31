package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var complete chan int = make(chan int)
var sep = flag.String("p", " ", "port")

func main() {
	// if len(os.Args) < 2 {   //os.Args 获致命令行参数
	// 	panic("没有命令行参数")
	// 	// os.Exit(1)
	// }

	//go run client.go -p 8000,8001
	flag.Parse()
	// fmt.Printf("%T,%[1]v", *sep) //
	for _, x := range strings.Split(*sep, ",") {
		go cli(x)
	}
	<-complete //阻塞 推出
}

func cli(x string) {
	conn, err := net.Dial("tcp", "192.168.1.105:"+x)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn) // 回声174页
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
