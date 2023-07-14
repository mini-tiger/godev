package main

import (
	"context"
	"fmt"
	"time"
)

type TaskFunc func(ctx context.Context) error

func DoTask(task TaskFunc, timeout time.Duration) error {
	// 创建一个具有设定超时时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 开始执行任务
	errChan := make(chan error)
	go func() {
		errChan <- task(ctx) //xxx  任务函数直接返回 err 即可
		close(errChan)
	}()

	// 等待任务结束
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout:", ctx.Err())
		return ctx.Err()
	}
}

func main() {

	// 测试任务
	err := DoTask(func(ctx context.Context) error {
		fmt.Println("Task started")
		time.Sleep(6 * time.Second)
		fmt.Println("Task completed")
		return nil
	}, 5*time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	}
	// 保证不退出
	finchan := make(chan struct{}, 0)
	go func() {
		time.Sleep(1000 * time.Second)
		finchan <- struct{}{}
	}()
	<-finchan
}
