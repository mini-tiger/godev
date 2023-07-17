package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func DoTask1(timeout time.Duration) error {
	defer wg.Done()
	// 创建一个具有设定超时时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 开始执行任务
	errChan := make(chan error)
	go func() {
		errChan <- Task1(ctx) //xxx  任务函数直接返回 err 即可
		close(errChan)
	}()

	// 等待任务结束,上层只能得到 先返回 的
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout:", ctx.Err())
		return ctx.Err()
	}
}

func Task1(ctx context.Context) error {
	fmt.Println("Task started")
	time.Sleep(6 * time.Second)
	fmt.Println("Task completed")
	return nil

}

func main() {
	wg.Add(1)
	// 测试任务
	go func() {
		err := DoTask1(5 * time.Second)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// 保证不退出
	wg.Wait()
}
