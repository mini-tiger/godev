package main

import (
	"fmt"
	"log"
	"time"
)

func test_dobule(i int) (r int) {
	return i * 2
}
func ps2_ex(fn string) func() {
	log.Printf("%s :%s\n", fn, "代码第一次执行")   // 2) 2018/04/20 17:25:41 ps2 :代码第一次执行
	return func() { log.Println("返回调用执行") } // 3)2018/04/20 17:25:47 返回调用执行
}

func main() {
	ts := time.Now()
	fmt.Printf("%T,%[1]v\n", ts) // 1) time.Time,2018-04-20 17:12:14.197 +0800 CST

	defer func() {
		fmt.Printf("%v\n", "123") // 4) 123
		if err := recover(); err != nil {
			fmt.Println(err) // 5) 这里的err其实就是panic传入的内容，lt 5
		}

	}()

	defer ps2_ex("ps2")()
	if test_dobule(2) <= 5 {
		panic("lt 5")
	}

	//panic 类似python raise 断言 生成错误
	// 必须使用 defer 以及 内置函数recover,捕获
	// 下面 代码不执行
	// 已执行过的defer 还是 先进后出 运行规则

	defer func() { fmt.Printf("%v\n", "124") }()
	log.Printf("%v\n", test_dobule(1))

}

/*
output:
C:\godev\basic>go run panic_recover.go
time.Time,2018-04-20 17:24:45.514 +0800 CST
2018/04/20 17:24:45 ps2 :代码第一次执行
2018/04/20 17:24:45 返回调用执行
123
lt 5

*/
