package main

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"
)

/*

https://www.cnblogs.com/golove/p/5909968.html


指针类型：

*类型：普通指针，用于传递对象地址，不能进行指针运算。

unsafe.Pointer：通用指针类型，用于转换不同类型的指针，不能进行指针运算,赋值

uintptr：用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

　　unsafe.Pointer 可以和 普通指针 进行相互转换。
　　unsafe.Pointer 可以和 uintptr 进行相互转换。

　　也就是说 todo unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。
*/
type A struct {
	AA int
}

func main() {
	var i int = 1
	fmt.Printf("i变量类型%T, 内容是:%d\n", i, i)

	p := unsafe.Pointer(&i) // todo 将i 的地址传入，转化为通用指针
	fmt.Printf("p变量类型%T, 内容是:%v\n", p, p)

	pu := (*int)(p) // 将通用指针转换为 int 型指针,不要转换为变量本身以外的类型
	*pu = 12
	fmt.Printf("i变量类型%T, 内容是:%d\n", i, i)

	print("===========================================")
	a := A{1}
	a1 := A{2}
	fmt.Printf("a:%v,addr:%p\n", a, &a)
	fmt.Printf("a1:%v,addr:%p\n", a1, &a1)

	au:=unsafe.Pointer(&a)
	au1:=unsafe.Pointer(&a1)
	a2:=atomic.SwapPointer(&au,au1)
	print("atomic.SwapPointer(&au,au1)转换后\n")
	fmt.Printf("a:%v,addr:%p\n", a, &a)
	fmt.Printf("a1:%v,addr:%p\n", a1, &a1)

	fmt.Printf("au:%v,addr:%p\n",(*A)(au), &au)
	fmt.Printf("au1:%v,addr:%p\n",(*A)(au1), &au1)
	fmt.Printf("a2Type:%v,a2:%v,addr:%p\n", reflect.TypeOf(a2),(*A)(a2) ,&a2)
}
