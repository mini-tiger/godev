package main

import (
	"context"
	"fmt"
	"time"
)

var exitChan chan struct{} = make(chan struct{}, 0)

func main() {
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d) // 50s 后

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.

	go func() {
		time.Sleep(10 * time.Second) //10s后强制退出
		cancel()
	}()

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("overslept")
			case <-ctx.Done():

				exitChan <- struct{}{}

				fmt.Println(ctx.Err())

				break
			}
		}

	}()

	<-exitChan
	fmt.Println("Exit")
}
