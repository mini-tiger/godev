package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  数组
 * @Version: 1.0.0
 * @Date: 2021/5/17 下午2:55
 */

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range a { //range是副本；循环次数在循环开始前就已经确定，循环内改变切⽚的⻓度，不影响循环次数
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // [1 2 3 4 5]
	fmt.Println("a = ", a) // [1 12 13 4 5]

	fmt.Println("以下是使用 数组指针")

	a = [5]int{1, 2, 3, 4, 5}
	r = [5]int{}
	for i, v := range &a { //使⽤ *[5]int 作为 range 表达式，其副本依旧是⼀个指向原数组 a 的指针，因此后续所有 循环中均是 &a 指向的原数组亲⾃参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // [1 12 13 4 5]
	fmt.Println("a = ", a) // [1 12 13 4 5]

	slice_range()
}

func slice_range() {
	fmt.Println("==========")
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range a { //range是副本；切片副本 同时指向同一个底层数组
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // [1 12 13 4 5]
	fmt.Println("a = ", a) // [1 12 13 4 5]
}
