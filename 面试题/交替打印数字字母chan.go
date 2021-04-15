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

var chan3 chan int = make(chan int, 0)
var chan4 chan string = make(chan string, 0)

//var num = 0
var str int32 = 'A'
var sw1 sync.WaitGroup

func main() {
	go func() {
		for {
			select {
			case t := <-chan3:
				fmt.Println(t)
				sw1.Done()
			case t := <-chan4:
				fmt.Println(t)
				sw1.Done()
			}
		}
	}()

	for i := 0; i < 26; i++ {

		sw1.Add(2)
		chan3 <- i
		chan4 <- string(str + int32(i))

	}

	sw1.Wait()

}
