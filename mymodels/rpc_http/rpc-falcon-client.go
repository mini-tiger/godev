package main

import (
	"time"
	"sync"
	"net/rpc"
	"log"
	"github.com/toolkits/net"
	"math"
	"math/rand"
	"godev/mymodels/rpc_http/common"
)

type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

var (
	HbsClient *SingleConnRpcClient
)

func (this *SingleConnRpcClient) close() {
	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

func (this *SingleConnRpcClient) serverConn() error {
	if this.rpcClient != nil {
		return nil
	}

	var err error
	var retry int = 1

	for {
		if this.rpcClient != nil {
			return nil
		}

		this.rpcClient, err = net.JsonRpcClient("tcp", this.RpcServer, this.Timeout)
		if err != nil {
			log.Printf("dial %s fail: %v", this.RpcServer, err)
			if retry > 3 {
				return err
			}
			time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)
			retry++
			continue
		}
		return err
	}
}

func (this *SingleConnRpcClient) Call(method string, args interface{}, reply interface{}) error {

	this.Lock()
	defer this.Unlock()

	err := this.serverConn()
	if err != nil {
		return err
	}

	timeout := time.Duration(10 * time.Second)
	done := make(chan error, 1)

	go func() {
		err := this.rpcClient.Call(method, args, reply)
		done <- err
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
		this.close()
	case err := <-done:
		if err != nil {
			this.close()
			return err
		}
	}

	return nil
}
func init() {
	HbsClient = &SingleConnRpcClient{
		RpcServer: "127.0.0.1:7777",
		Timeout:   time.Duration(2000) * time.Millisecond,
	}
}

func main() {
	f:=func() {
		for {
			req := common.Recv{}
			req.Num1 = rand.Intn(100)
			req.Num2 = rand.Intn(50)
			var resp common.ResultResp
			err := HbsClient.Call("Falcon.Compute", req, &resp)
			if err != nil  {
				log.Println("call Falcon.compute fail:", err, "Request:", req, "Response:", resp)
			}
			log.Println("call Falcon.compute Result:",resp.Result)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
	go f()
	go f()
	go f()
	select {}
}
