package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  切片
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午9:10
 */

func func1(num ...int) {
	fmt.Printf("%p\n", num) // 切片第一个元素地址
	num[0] = 18
	fmt.Printf("%p\n", &num[0]) // 切片中第一个元素地址
}

func init_slice() {
	nn := []int{2: 2, 3, 0: 1} // 2:2 索引2时   3  前面索引加1也就是索引3是3， 0:1 索引0是1 ,  缺少索引1 补0
	fmt.Println(nn)            // [1 0 2 3]

	//
	var nn1 []int
	nn1 = append(nn1, 1)
	fmt.Println(nn1) // 没 make 可以append
}

func main() {
	n := []int{1, 2, 3}
	fmt.Printf("%p\n", n) // 切片第一个元素地址
	func1(n...)
	fmt.Printf("%+v\n", n)

	init_slice()

	func2()

	func3()

	nilslice()

	loop()
}

func func2() {
	println(".....func2......")
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0], len(t), cap(t))

	s1 := a[:0]
	s2 := a[:2]
	s3 := a[1:2:cap(a)]
	fmt.Println(len(s1), cap(s1)) // 0,5   cap 随底层数组
	fmt.Println(len(s2), cap(s2)) // 2,5
	fmt.Println(len(s3), cap(s3)) // 1,4

}

/*
知识点：操作符 [i, j] 。基于数组（切⽚）可以使⽤操作符 [i, j] 创建新的切⽚，从索引 i
，到索引 i ，到索引 j 结束，截取已有数组（切⽚）的任意部分，返回新的切⽚，新切⽚的值包
含原数组（切⽚）的 i 索引的值，但是不包含 j 索引的值。 i 、 j 都是可选的， i 如果省略，
默认是0， j 如果省略，默认是原数组（切⽚）的⻓度。 i 、 j 都不能超过这个⻓度值。
假如底层数组的⼤⼩为 k，截取之后获得的切⽚的⻓度和容量的计算⽅法：⻓度：j-i，容ᰁ：k-i。
截取操作符还可以有第三个参数，
xxx 形如 [i,j,k]，第三个参数 k ⽤来限制新切⽚的容量，但不能超过 原数组（切⽚）的底层数组⼤⼩。截取获得的切⽚的⻓度和容量分别是：j-i、k-i
以例⼦中，切⽚ t 为 [4]，⻓度和容量都是 1。
*/

func func3() {
	println(".....func3......")
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal") // 必须是长度 元素顺序完全栗
	} else {
		fmt.Println("not equal")
	}

	//b=append(b,7)
	//if a == b {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
	/*
		xxx 切片不能比较
		Go中的数组是值类型，可⽐较，另外⼀⽅⾯，数组的⻓度也是数组类型的组成部分，所以 a 和 b 是不
		同的类型，是不能⽐较的，所以编译错误
	*/
}
func nilslice() {
	println("......nilslice.....")
	var s []int
	var ss []int = []int{}
	fmt.Println(s == nil)  // true
	fmt.Println(ss == nil) //false  ；空切⽚和 nil 不相等，表示⼀个空的集合
}

func loop() {
	println("......loop.....")
	v := []int{1, 2, 3}
	for i, _ := range v { // range是副本；循环次数在循环开始前就已经确定，循环内改变切⽚的⻓度，不影响循环次数
		v = append(v, i)
		fmt.Printf("cureent slice:%+v\n", v)
	}

}
