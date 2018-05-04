package main

import (
	"fmt"
)

func main() {
	// var i *int

	i := new(int) //分配内存地址,返回一个指向该类型内存地址的指针。同时请注意它同时把分配的内存置为零，也就是类型的零值。

	// *i = 10
	*i = 10
	fmt.Println(*i) //10

	// var p *[]int = new([]int)
	// 或
	p := new([]int) //不能直接使用

	// var v []int = make([]int, 0)
	v := make([]int, 0)
	fmt.Println(*p, p, v) //[] &[] []

	vv := make([]int, 2)
	vv[0] = 1
	fmt.Println(vv)

	var x []int

	fmt.Println(x)
}
