package main

import "fmt"

/*
http://www.runoob.com/go/go-data-types.html

*/

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

func init() { // 最先执行的 函数

	var b string
	b = "this is 注释"
	fmt.Println(b)
	fmt.Print(b, "\n")
}

func main() {
	var v_name, v_type = 1, 2
	fmt.Println(v_name, v_type)
	var_ex()
}

func var_ex() {
	fmt.Println(" this is variables ")

	var u int32
	u = 128
	fmt.Printf("u: %d \n", u)

	var m int = 10
	fmt.Printf("m: %d \n", m)

	var i, j, k, s = 1, 2, 3, "diy" //自动判断 类型
	//注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	fmt.Printf("i: %d ,j: %d ,k: %d,s: %s \n", i, j, k, s)

	var (
		num  int
		name string
	)
	num = 1
	name = "zhangsan"
	fmt.Printf("num: %d ,name %s \n", num, name)

	var b = true //bool型
	fmt.Printf("bool: %t \n", b)
	// var bb int

	num, h := 123, "hello" //这种不带声明格式的只能在函数体中出现
	fmt.Printf("g: %d ,h %s \n", num, h)

	_, _aa, Ret := 1, 2, 3 //Go中有一个特殊的变量_ 任何赋给它的值将被丢弃
	fmt.Printf("_aa：%d ,Ret: %d\n", _aa, Ret)

	var a = 1
	var aa int
	var bb = 1

	aa = a
	fmt.Println(&a, &aa, &bb)
	fmt.Println(a, aa, bb)
}
