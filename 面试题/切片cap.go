package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  切片cap
 * @Version: 1.0.0
 * @Date: 2021/5/14 下午3:30
 */

func main() {
	firstslice := []int{1, 2, 3}
	fmt.Printf("%+v\n", firstslice) // [1 2 3]
	ap(firstslice)
	fmt.Printf("%+v\n", firstslice) // [10 2 3]
	app(firstslice)                 // 超过源slice cap ,
	fmt.Printf("%+v\n", firstslice) // [10 2 3]

}
func app(f []int) {
	f = append(f, 4)
	/*
		因为append导致底层数组重新分配内存了，append中的f这个slice的底层数组和外⾯不是⼀个，并没
		有改变外⾯的。
	*/
}
func ap(f []int) {
	f[0] = 10
}
