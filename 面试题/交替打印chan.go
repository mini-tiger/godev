package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  交替打印
 * @Version: 1.0.0
 * @Date: 2021/4/15 下午3:44
 */

var chan1 chan int = make(chan int, 0)
var chan2 chan int = make(chan int, 0)
var sw sync.WaitGroup

func main() {
	go func() {
		for {
			select {
			case num := <-chan1:
				fmt.Println(num)
				sw.Done()
			case num := <-chan2:
				fmt.Println(num)
				sw.Done()
			}
		}
	}()

	for i := 0; i < 100; i++ {
		sw.Add(1)
		if i%2 == 0 {
			chan1 <- i
		} else {
			chan2 <- i
		}
	}

	sw.Wait()

}
