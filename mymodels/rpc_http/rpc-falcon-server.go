package main

import (
	"net/rpc"
	"net"
	"log"
	"time"
	"net/rpc/jsonrpc"
	"godev/mymodels/rpc_http/common"
)

func Start() {
	//addr := g.Config().Listen
	addr := "127.0.0.1:7777"
	server := rpc.NewServer()
	// server.Register(new(filter.Filter))
	server.Register(new(common.Falcon))

	//server.Register(new(Port))

	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("listener accept fail:", err)
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

func main()  {
	go Start()
	select {}
}