package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "goroutine1")
	go watch(ctx, "goroutine2")

	time.Sleep(3 * time.Second)
	fmt.Println("可以了，通知 all goroutine停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println("监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
