package main

import (
	"fmt"
	"log"
	"time"
)

func long(aa int) (r int) {

	// fmt.Printf("%T,%[1]v\n", aa) //[]int,[1 2 2]
	defer func() { r += aa }() //2
	//注意最后小括号，在rerurn 之后执行
	//匿名函数可以 获取 外层函数 变量
	if aa > 2 {
		defer func() { aa += 1 }() //no run
	}

	return aa
}

func ps2_ex(fn string) func() {
	log.Printf("%s :%s\n", fn, "代码第一次执行")
	return func() { log.Println("返回调用执行") }
}

func ps2() (r int) {
	// start := time.Now()
	log.Printf("enter %s", "ps2") //最先打印
	time.Sleep(1 * time.Second)
	defer ps2_ex("ps2")()
	//注意最后括号，否则不执行 return func() { log.Println("返回调用执行") }
	//等待上面的1秒  log.Printf("%s :%s\n", fn, "代码第一次执行")
	//()也说是返回的匿名funv, 在等待下面1秒后执行 return func() { log.Println("返回调用执行") }

	time.Sleep(1 * time.Second)
	/*
		2018/04/19 21:26:26 enter ps2
		2018/04/19 21:26:27 ps2 :代码第一次执行
		2018/04/19 21:26:28 返回调用执行
	*/
	return 11

}

func main() {
	fmt.Println(ps1(1))

	/*
	   11
	   12
	   2
	   2
	*/
	a := 1
	fmt.Println(long(a)) //5 ,
	fmt.Println(ps2())   //ps2()

	/*
		2018/04/19 21:26:26 enter ps2
		2018/04/19 21:26:27 ps2 :代码第一次执行
		2018/04/19 21:26:28 返回调用执行
	*/
}

func ps1(s int) (result int) {

	// defer fmt.Printf("%s\n", result+2)
	defer fmt.Println(12) //多个defer 执行顺序是从后到前
	defer fmt.Println(11)
	return s + 1
	/*
	   11
	   12
	   2
	*/
}
