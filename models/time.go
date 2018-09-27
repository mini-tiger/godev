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
	fmt.Println(time.Now().UnixNano())
	seconds := 10
	fmt.Print(time.Duration(seconds)*time.Minute, "\n") // 打印 10m0s,单位time.Duration(seconds)
	d, _ := time.ParseDuration("3m10s")                 // 合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"。
	fmt.Printf("%T,%[1]v\n", d)
	fmt.Println(d.Seconds())
	time_start := time.Now()
	//time.NewTicker(d) 每过d秒后触发，返回通道字段
	chan1 := time.NewTicker(1 * time.Second)
	for i := 0; i < 2; i++ {
		fmt.Println(<-chan1.C)
	}
	// fmt.Printf("%d\n", 1)
	//fmt.Println(<-chan1.C) //提取时间

	//时间格式
	t1 := time.Now()
	fmt.Println(t1.Format(time.ANSIC)) //必须使用 文档中定义的时间不能修改
	fmt.Println(t1.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t1.Format("Mon Jan _2 15:04:06 2006")) //todo 修改后，输出时间倒退了
	fmt.Println(t1.Format("2006 01-02 15:04:05"))
	fmt.Println(t1.Year(), int(t1.Month()), t1.Day(), t1.Hour(), t1.Minute(), t1.Second()) //todo 自定义 2018 9 6 16 6 33
	fmt.Println(t1.Date())
	fmt.Println(int(time.September))
	fmt.Printf("过去了 :%s 秒\n", time.Since(time_start))

	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	logDay := yesTime.Format("20060102")
	fmt.Printf("%T,%v\n", logDay, logDay) // todo 昨天日期

	base_format := "2006-01-02 15:04:05"
	str_time := nTime.Format(base_format)
	parse_str_time, _ := time.Parse(base_format, str_time) // todo 字符串转时间
	fmt.Printf("string to datetime :%v\n", parse_str_time)

	base_format = "2006-01-02 15:04"
	s := "2018-09-24 17:23"                         //todo 注意月份 日期 是否 有0，09还是0
	parse_str_time, _ = time.Parse(base_format, s) // todo 字符串转时间
	fmt.Printf("string to datetime :%s\n", parse_str_time)
	fmt.Println(7*86400 + parse_str_time.Unix())
}