package main

import (
	"fmt"
	"unsafe"
)

//对齐  pages 278
//只有 bool  intN floatN 不一定是8字节整数倍
func main() {

	fmt.Println(unsafe.Sizeof(aaa))
	type name struct {
		// b bool
		// x float64
		y []int
		x float32
		b bool
	}
	var (
		n name
	)

	fmt.Println(unsafe.Sizeof(n.y), unsafe.Sizeof(n.x), unsafe.Sizeof(n.b)) // 1 4 24
	//对齐 8的整数按8 算，一个字节
	fmt.Println(unsafe.Alignof(n.y), unsafe.Alignof(n.x), unsafe.Alignof(n.b)) //8 4 1
	fmt.Println(unsafe.Sizeof(n))                                              //32
	/*
			32 位系统4字节一个字，64位8字节一个字

		不同类型 占用空间不一样
		b bool 1个字节
		x floatN  N*8 个字节


		[]T  24个字节  3个字，字长度8

		name struct
		x 占4
		b 占1
		由于8字节是1个字，x和b 挨着可以一起占用一个字

	*/
	fmt.Println("-------------------------------------")
	ex1()

	fmt.Println("-------------------------------------")
	ex2()
}
func ex1() {

	type name struct {
		b bool
		a float32
		y map[string]int
		x interface{}
		e *int
		// b bool
	}
	var (
		n name
	)
	fmt.Println(unsafe.Sizeof(n.y), unsafe.Sizeof(n.x), unsafe.Sizeof(n.b))    // 16 16 1
	fmt.Println(unsafe.Alignof(n.y), unsafe.Alignof(n.x), unsafe.Alignof(n.b)) //8 8 1

	fmt.Println(unsafe.Sizeof(n))
	/*
		//只有 bool  intN floatN 不一定是8字节整数倍

		现在是32

		将b放到下面则是40
		因为a,b 可以共用8字节，分开则要分别占用8字节

	*/
}
func ex2() {
	type name struct {
		b bool           // 偏移量开头0
		a float32        //占用4，前面是BOOL，占用同一个8字节（1个字）， 开头 0+4
		y map[string]int //上面两个为一行占用8字节，新起一行  8+0
		x interface{}    //上面map占用一个字，8+8
		e *int           //上面interface占用2个字，8+8+16
		d bool           //上面 指针用2个字，8+8+16+8
	}
	var (
		n name
	)

	//偏移量第一个0, 依次加上 内存占用量
	fmt.Println(unsafe.Offsetof(n.b), unsafe.Offsetof(n.a), unsafe.Offsetof(n.y),
		unsafe.Offsetof(n.x), unsafe.Offsetof(n.e), unsafe.Offsetof(n.d))
}
