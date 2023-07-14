package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go doSomething(ctxWithTimeout, done)

	time.Sleep(4 * time.Second)

	cancel()
	<-done
}

func doSomething(ctx context.Context, done chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			// 执行清理操作
			fmt.Println("ctx.Done...", ctx.Err())
			done <- struct{}{}
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Doing something...")
		}
	}
}
