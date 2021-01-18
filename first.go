package main

import "fmt"

func fun() *int { //int类型指针函数
	var tmp int = 1
	defer fmt.Println(tmp * 2)
	tmp = 2
	return &tmp //返回局部变量tmp的地址
}

func main() {
	var p *int
	p = fun()
	fmt.Printf("%d\n", *p) //这里不会像C，报错段错误提示，而是成功返回变量V的值1
}
