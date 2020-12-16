package main

// rpc  传送list可能会有问题
import (
	"godev/mymodels/rpc_http/common"
	"log"
	"net"
	"net/rpc"
	"time"
	//_ "github.com/CodyGuo/godaemon"
)

func main() {
	//方式一
	var ms = new(common.MathService)
	//rpc.Register(ms)
	//rpc.HandleHTTP() //将Rpc绑定到HTTP协议上。
	//fmt.Println("启动服务...")
	//err := http.ListenAndServe(":1234", nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("服务已停止!")

	//方式二
	newServer := rpc.NewServer()
	newServer.Register(ms)
	ll, e := net.Listen("tcp", "127.0.0.1:1234") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	newServer.HandleHTTP("/ext", "/bar")
	for {
		go newServer.Accept(ll)

		time.Sleep(2 * time.Second)
	}

}
