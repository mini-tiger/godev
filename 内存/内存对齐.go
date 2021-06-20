package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// https://xie.infoq.cn/article/594a7f54c639accb53796cfc7
// https://go.wuhaolin.cn/gopl/ch13/ch13-01.html
// xxx 对齐从 占用 小 到大 排序，如果达到最大对齐8（占用大于等于8），就无法优化

type Demo2 struct {
	S  [1]int8 // 1字节
	A2 int64   // 8字节
	A3 int16   // 2字节

}

type Demo3 struct {
	S  [1]int8
	A3 int16
	A2 int64
}

type More struct {
	c  chan int8      // chan 1*8 =8
	m  map[int]string // map 1*8 =8
	s  string         // 占用 2*8=16  data len
	ss []int          // 占用 3*8=24  data len  cap
	i  interface{}    // 2*8=16 type,value
}

func main() {

	fmt.Printf("Demo2 占用内存: %d\n", unsafe.Sizeof(Demo2{}))

	fmt.Printf("Demo2 S type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo2{}.S),
		unsafe.Offsetof(Demo2{}.S),
		unsafe.Sizeof(Demo2{}.S),
		unsafe.Alignof(Demo2{}.S))

	fmt.Printf("Demo2 A2 type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo2{}.A2),
		unsafe.Offsetof(Demo2{}.A2),
		unsafe.Sizeof(Demo2{}.A2),
		unsafe.Alignof(Demo2{}.A2))

	fmt.Printf("Demo2 A3 type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo2{}.A3),
		unsafe.Offsetof(Demo2{}.A3),
		unsafe.Sizeof(Demo2{}.A3),
		unsafe.Alignof(Demo2{}.A3))

	/*
		xxx 最大对齐是 8
		Demo2 :
		S 占用1 , A2 占用8, A3 占用2
		a代表有 x代表填充
		axxx xxxx aaaa aaaa aaxx xxxx

		S占用2  从头开始占用1个
		A2占用8  前面s填充7个，从偏移量8开始 写 8个
		A3占用2  前面A3不是填充，从偏移量16开始 写 2个，在填充6个

		Demo3 :
		S 占用1 , A3 占用2 , A2 占用8
		a代表有 x代表填充
		axaa xxxx aaaa aaaa

		S占用2  从头开始占用1个
		A3占用2  填充1个，从偏移量2开始 写 2个
		A2占用8  前面A3填充4个，从偏移量8开始 写 8个

	*/
	fmt.Println("重排优化")

	fmt.Printf("Demo3 占用内存: %d\n", unsafe.Sizeof(Demo3{}))

	fmt.Printf("Demo3 S type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo3{}.S),
		unsafe.Offsetof(Demo3{}.S),
		unsafe.Sizeof(Demo3{}.S),
		unsafe.Alignof(Demo3{}.S))

	fmt.Printf("Demo3 A3 type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo3{}.A3),
		unsafe.Offsetof(Demo3{}.A3),
		unsafe.Sizeof(Demo3{}.A3),
		unsafe.Alignof(Demo3{}.A3))

	fmt.Printf("Demo3 A2 type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(Demo3{}.A2),
		unsafe.Offsetof(Demo3{}.A2),
		unsafe.Sizeof(Demo3{}.A2),
		unsafe.Alignof(Demo3{}.A2))

	fmt.Println("其它占用")

	fmt.Printf("More 占用内存: %d\n", unsafe.Sizeof(More{}))

	fmt.Printf("More c type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(More{}.c),
		unsafe.Offsetof(More{}.c),
		unsafe.Sizeof(More{}.c),
		unsafe.Alignof(More{}.c))

	fmt.Printf("More{} m type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(More{}.m),
		unsafe.Offsetof(More{}.m),
		unsafe.Sizeof(More{}.m),
		unsafe.Alignof(More{}.m))

	fmt.Printf("More{} s type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(More{}.s),
		unsafe.Offsetof(More{}.s),
		unsafe.Sizeof(More{}.s),
		unsafe.Alignof(More{}.s))

	fmt.Printf("More{} ss type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(More{}.ss),
		unsafe.Offsetof(More{}.ss),
		unsafe.Sizeof(More{}.ss),
		unsafe.Alignof(More{}.ss))

	fmt.Printf("More{} i type: %v ,偏移: %d , 占用: %d ,对齐: %d \n",
		reflect.TypeOf(More{}.i),
		unsafe.Offsetof(More{}.i),
		unsafe.Sizeof(More{}.i),
		unsafe.Alignof(More{}.i))
}
