package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cb := make(chan int, 1)

	go func() {
		fmt.Printf("%v\n", <-cb) //有缓存的通信，不阻塞，所以不输出
	}()
	cb <- 1

	go func() {
		fmt.Printf("%v\n", <-c) //阻塞
	}()
	c <- 11
}
