package main

import (
	"net"
	"log"
)

func serverConn(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
		//	if err == io.EOF {
		//		log.Println("server io EOF\n")
		//		return
		//	}
			log.Printf("[ERROR] conn:%s,server read faild: %s\n", conn.RemoteAddr(),err)
			return
		}
		log.Printf("recevice %d bytes, content is 【%s】\n", n, string(buf[:n]))


	}
}

func main() {
	// 建立监听
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("error listen: %s\n", err)
	}
	defer l.Close()
	for {

		log.Println("waiting accept.")
		// 允许客户端连接，在没有客户端连接时，会一直阻塞
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("accept faild: %s\n", err)
		}
		serverConn(conn)
	}
}
