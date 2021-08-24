package main

import (
	"context"
	"fmt"
	"time"
)

var key1 string = "name1"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx1 := context.WithValue(ctx, key1, "[goroutine1]")

	go watch(valueCtx1)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知 all goroutine停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println("监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key1), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
