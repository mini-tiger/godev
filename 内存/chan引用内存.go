package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch2 := ch1
	ch1 <- 123
	ch2 <- 456
	fmt.Println(ch1, ch2)     //输出地址一样
	fmt.Println(<-ch2, <-ch2) // 指向内存一样，从哪个通道 获取都 可以
}
