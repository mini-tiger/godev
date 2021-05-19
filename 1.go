package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  结构体
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午11:26
 */

type N int

func (n N) test() {
	fmt.Println(n)
}
func main() {
	var n N = 10
	fmt.Println(n)
	n++
	f1 := N.test
	f1(n)
	n++
	f2 := (*N).test
	f2(&n)
}
