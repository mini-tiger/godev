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

func ex() {
	fs := [3]func(){}
	// fmt.Println(fs)
	for i := 0; i < len(fs); i++ {
		defer fmt.Println("defer i:", i) // 直接引用i ,则是i 每次值的 复制
		defer func() {                   //在func内，引用外层 i变量的内存地址，i在for中被修改多次，最后会是3,3不满足i<len(fs) ,不执行坛坛循环体
			fmt.Println("defer func i:", i)
		}()
		fs[i] = func() { fmt.Println("func fs i:", i) } //这里也是func内

	}
	for _, f := range fs {
		f()
	}

}

func ex1() {
	fs := [3]func(){}
	// fmt.Println(fs)
	for i := 0; i < len(fs); i++ {
		defer fmt.Println("defer i:", i) // 直接引用i ,则是i 每次值的 复制
		defer func(i int) {
			fmt.Println("defer func i:", i)
		}(i)
		fs[i] = func() { fmt.Println("func fs i:", i) } //这里也是func内

	}
	for _, f := range fs {
		f()
	}

}

func main() {
	ex()
	fmt.Println("-------------------------------------")
	ex1()

	/*
		func fs i: 4
		func fs i: 4
		func fs i: 4
		func fs i: 4
		defer func i: 4
		defer i: 3
		defer func i: 4
		defer i: 2
		defer func i: 4
		defer i: 1
		defer func i: 4
		defer i: 0
	*/
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
