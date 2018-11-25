package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Printf("i变量类型%T, 内容是:%d\n", i, i)

	p := unsafe.Pointer(&i) // todo 将i 的地址传入，转化为通用指针
	fmt.Printf("p变量类型%T, 内容是:%v\n", p, p)

	pu := (*int)(p)
	*pu = 12
	fmt.Printf("i变量类型%T, 内容是:%d\n", i, i)

}
