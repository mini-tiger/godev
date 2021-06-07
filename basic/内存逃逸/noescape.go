package main

import (
	"unsafe"
)

type A struct {
	S *string
}

func (f *A) String() string {
	return *f.S
}

type ATrick struct {
	S unsafe.Pointer
}

func (f *ATrick) String() string {
	return *((*string)(f.S))
}

func NewA(s string) A { //  moved to heap: s
	return A{S: &s}
}

func NewATrick(s string) ATrick {
	return ATrick{S: noescape(unsafe.Pointer(&s))} //s does not escape
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func main() {
	s := "hello"
	f1 := NewA(s)
	f2 := NewATrick(s)
	s1 := f1.String()
	s2 := f2.String()
	_ = s1 + s2
}

//上段代码对A和ATrick同样的功能有两种实现：他们包含一个 string ，然后用 String() 方法返回这个字符串。但是从逃逸分析看ATrick 版本没有逃逸。
//
//noescape() 函数的作用是遮蔽输入和输出的依赖关系。使编译器不认为 p 会通过 x 逃逸， 因为 uintptr() 产生的引用是编译器无法理解的。
//
//内置的 uintptr 类型是一个真正的指针类型，但是在编译器层面，它只是一个存储一个 指针地址 的 int 类型。代码的最后一行返回 unsafe.Pointer 也是一个 int。
//
//noescape() 在 runtime 包中使用 unsafe.Pointer 的地方被大量使用。如果作者清楚被 unsafe.Pointer 引用的数据肯定不会被逃逸，但编译器却不知道的情况下，这是很有用的。
