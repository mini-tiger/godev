package main

import (
	"./p"
	"fmt"
	"unsafe"
)

func main() {
	var v p.V
	// v.i = 32 //由于 i小写不能导出
	fmt.Println(v)
	ex()
}

func ex() {
	//https://studygolang.com/articles/1414

	var v *p.V = new(p.V) //v就是类型为p.V的一个指针

	var i *int32 = (*int32)(unsafe.Pointer(v)) //将指针v转成通用指针，再转成int32指针。

	*i = int32(98)

	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	//i是int32类型，也就是说i占4个字节。所以j是相对于v偏移了4个字节。您可以用uintptr(4)或uintptr(unsafe.Sizeof(int32(0)))来做这个事。unsafe.Sizeof方法用来得到一个值应该占用多少个字节空间
	*j = int64(763)

	v.PutI()

	v.PutJ()

}
