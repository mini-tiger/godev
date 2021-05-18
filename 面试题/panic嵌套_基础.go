package main

import (
	"fmt"
	"time"
)

func main() {
	defer println("in main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test", err)
		}
	}()
	go func() { // xxx 在gorouine内 外面recover() 捕捉不到
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(1 * time.Second)
}

// xxx panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
// recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数，在其他作用域中调用不会发挥作用；
// xxx panic 只会触发当前 Goroutine 的 defer
