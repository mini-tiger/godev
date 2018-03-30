package main

import "fmt"

/*
this is 注释
*/

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

func init() {

	var b string
	b = "this is 注释"
	fmt.Println(b)
	fmt.Print(b, "\n")
}

func main() {
	var v_name, v_type = 1, 2
	fmt.Println(v_name, v_type)
}
