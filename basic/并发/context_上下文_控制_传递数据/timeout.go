package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg_timeout *log.Logger

func someHandler_timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second)) // xxx 5s超时 ctx.Done <-
	go doStuff_timeout(ctx)

	//10秒后取消doStuff
	time.Sleep(2 * time.Second)
	cancel() //xxx  调用 直接 ctx.Done <-

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff_timeout(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg_timeout.Printf("done")
			return
		default:
			logg_timeout.Printf("work")
		}
	}
}

func main() {
	logg_timeout = log.New(os.Stdout, "", log.Ltime)
	someHandler_timeout()
	logg_timeout.Printf("down")
}
