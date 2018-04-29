package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	// "os"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func listen(x string) {
	fmt.Println(x)
	listener, err := net.Listen("tcp", "localhost:"+x)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

var complete chan int = make(chan int)
var sep = flag.String("p", " ", "port")

func main() {
	// if len(os.Args) < 2 {   //os.Args 获致命令行参数
	// 	panic("没有命令行参数")
	// 	// os.Exit(1)
	// }

	//go run server.go -p 8000,8001
	flag.Parse()
	// fmt.Printf("%T,%[1]v", *sep) //
	fmt.Println("listen port:")
	for _, x := range strings.Split(*sep, ",") {
		go listen(x)
	}
	<-complete //阻塞 推出
}
