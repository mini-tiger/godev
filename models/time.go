package main

import (
	"fmt"
	"log"
	"time"
)

//!+main
func bigSlowOperation() int {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(2 * time.Second) // simulate slow operation by sleeping
	return 1
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	log.Println(bigSlowOperation())

	seconds := 10
	fmt.Print(time.Duration(seconds)*time.Minute, "\n") // 打印 10m0s,单位time.Duration(seconds)
	d, _ := time.ParseDuration("3m10s")
	fmt.Printf("%T,%[1]v\n", d)
	fmt.Println(d.Seconds())
	//time.NewTicker(d) 每过d秒后触发，返回当前时间到通道,
	chan1 := time.NewTicker(1 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println(<-chan1.C)
	}
	// fmt.Printf("%d\n", 1)
	fmt.Println(<-chan1.C) //提取时间

	//时间格式
	t1 := time.Now()
	fmt.Println(t1.Format(time.ANSIC)) //必须使用 文档中定义的时间不能修改
	fmt.Println(t1.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t1.Format("Mon Jan _2 15:04:06 2006")) // 修改后，输出时间倒退了
	fmt.Println(t1.Format("2006 01-02 15:04:05"))
}
