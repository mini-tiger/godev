package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {


	// 整形5 的地址 强制转换为float
	num:=5
	floatnum :=(*float32)(unsafe.Pointer(&num))
	fmt.Println(floatnum)
	fmt.Println(*floatnum)


	str:="A"
	str1:="B"
	//var pstr,pstr1 unsafe.Pointer
	//var pstr1 unsafe.Pointer
	pstr:=unsafe.Pointer(&str)
	pstr1:=unsafe.Pointer(&str1)
	fmt.Printf("转换前s的值:%s,addr:%p\n",str,&str)
	fmt.Printf("转换前s1的值:%s,addr:%p\n",str1,&str1)
	fmt.Printf("转换前pstr的值:%v,addr:%p\n",pstr,&pstr)
	fmt.Printf("转换前pstr1的值:%v,addr:%p\n",pstr1,&pstr1)
	//(string(pstr))
	fmt.Println("=========================================================")
	atomic.StorePointer(&pstr,pstr1) // pstr1 中的值地址 存储到 ＊pstr
	fmt.Printf("转换后s的值:%s,无变化addr:%p\n",str,&str)
	fmt.Printf("转换后s1的值:%s,无变化addr:%p\n",str1,&str1)
	fmt.Printf("转换后pstr的值:%v,addr:%p\n",*(*string)(pstr),&pstr)
	fmt.Printf("转换后pstr1的值:%v,addr:%p\n",*(*string)(pstr1),&pstr1)


	// xxx 通过偏移量 修改 结构体
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用（偏移量为 0）
	up0 := uintptr(p)
	// 将通用指针转换为 byte 型指针
	pb := (*byte)(p)
	// 给转换后的指针赋值
	*pb = 10
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 2 个字段
	up := up0 + unsafe.Offsetof(s.b)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 3 个字段
	up = up0 + unsafe.Offsetof(s.c)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 30
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 4 个字段
	up = up0 + unsafe.Offsetof(s.d)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 int64 型指针
	pi := (*int64)(p)
	// 给转换后的指针赋值
	*pi = 40
	// 结构体内容跟着改变
	fmt.Println(s)
}
