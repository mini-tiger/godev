package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	var i int = 1
	p := unsafe.Pointer(&i) //转换为通用指针
	var ii int = 2
	p1 := unsafe.Pointer(&ii)

	fmt.Printf("%p\n", p) // 原本的地址
	fmt.Printf("%p\n", p1)

	pp := atomic.SwapPointer(&p, p1) // 交换变量地址，  pp是返回 p的地址

	fmt.Printf("%p\n", p) // 和P1地址一样
	fmt.Printf("%p\n", pp)
	fmt.Println(*(*int)(p)) // (*int)(p) 转换为整形指针  ,*(*int)(p)取指针的值
}
