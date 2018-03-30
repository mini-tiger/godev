package main

import "fmt"

func getSequence(s int) func() int {
	i := s
	// fmt.Println(s, i)
	return func() int {
		i += 1
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */

	nextNumber := getSequence(1)
	fmt.Println(nextNumber())
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	// fmt.Println(nextNumber())
	// fmt.Println(nextNumber())
	// fmt.Println(nextNumber())

	//  创建新的函数 nextNumber1，并查看结果
	// nextNumber1 := getSequence()
	// fmt.Println(nextNumber1())
	// fmt.Println(nextNumber1())
}
