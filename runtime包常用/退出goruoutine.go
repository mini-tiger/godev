package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  退出goruoutine
 * @Version: 1.0.0
 * @Date: 2021/4/15 下午5:37
 */

func main() {

	go func() {
		fmt.Println("abc")
		fun()
		fmt.Println("bcd") // 这里不打印
	}()

	time.Sleep(3 * time.Second)
}
func fun() {
	defer fmt.Println("fun") // 这里打印
	runtime.Goexit()         // xxx  结束当前goroutine
}
