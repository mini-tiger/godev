package main

import (
	"fmt"
	"reflect"
)

func foo() {
	defer fmt.Println("A") // xxx 6.打印这里
	fmt.Println("C")       // xxx 1. 先打印这里
	defer func() {         // 3. defer后进先出 ,
		defer func() {
			fmt.Println("D") // xxx 5. 打印这里，之后结束当前Goroutine
		}()
		panic(2) // 4.执行这里 触发当前Goroutine 下的defer， xxx 这是后面触发的panic
	}()
	panic(54321) // 2. 运行到这里后 开始触发 执行当前Goroutine 下的 defer
}

func Panic() {
	defer func() {
		pv := recover()
		fmt.Println("recover err :")
		fmt.Println("11:", pv, reflect.TypeOf(pv)) // 8.   pv 收集的是最后一次 panic 的2, 如果没有recover 两次panic都打印
	}()

	foo() // 7. 这里触发了panic 不会在执行下面的代码，触发本Goroutine 下的defer
	fmt.Println("end")
}

func main() {
	Panic()
}
