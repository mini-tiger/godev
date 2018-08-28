package main
// rpc  传送list可能会有问题
import (
	"fmt"
	"godev/mymodels/rpc_http/common"
	"net/rpc"
	"net"
	"log"
	"time"
)

func main() {
	//var args= common.Args{17, 8}
	//var result= common.Result{}
	//方式一
	//var client, err= rpc.DialHTTP("tcp", "127.0.0.1:5555")
	//
	//if err != nil {
	//	fmt.Println("连接RPC服务失败：", err)
	//}
	////method 是server端已经注册过的服务，client这边只需要写成字符串
	//err = client.Call("MathService.Divide", args, &result)
	//if err != nil {
	//	fmt.Println("调用失败：", err)
	//}
	//fmt.Println("调用结果：", result.Value)
	//方式二
	var args= common.Args{17, 8}
	var result= common.Result{}
	for {
		address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
		if err != nil {
			time.Sleep(1*time.Second)
			log.Println("1s after reconn")
			continue
		}
		conn, err := net.DialTCP("tcp", nil, address)
		if err!=nil {
			time.Sleep(1*time.Second)
			log.Println("1s after reconn")
			continue
		}
		defer conn.Close()
		client := rpc.NewClient(conn)
		defer client.Close()

		//reply := make([]string, 10) //调用服务端的方法，发送的args ，list会有问题
		err = client.Call("MathService.Divide", args, &result)
		if err != nil {
			fmt.Println("arith error:", err)
		}
		log.Println(result)
		time.Sleep(2*time.Second)
	}
}