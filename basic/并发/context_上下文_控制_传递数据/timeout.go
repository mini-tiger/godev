package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg_timeout *log.Logger

func someHandler_timeout() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second)) // xxx 5s超时 ctx.Done <-

	finchan := make(chan struct{}, 0)
	go func() {
		defer func() {
			finchan <- struct{}{}
		}()
		doStuff_timeout()

	}()

	for {
		select {
		case <-ctx.Done():
			logg_timeout.Printf("timeout %v", ctx.Err())
			return
		case <-finchan:
			logg_timeout.Printf("finish")
		}
	}

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff_timeout() {
	for {
		time.Sleep(1 * time.Second)
		logg_timeout.Printf("%v\n", time.Now())
	}

}

func main() {
	logg_timeout = log.New(os.Stdout, "", log.Ltime)
	someHandler_timeout()
	//logg_timeout.Printf("down")
}
