package main

import "fmt"

func testsub() (i int) {
	defer func() {
		i = 1
	}()
	panic(i) // 触发后 执行 当前函数内的 defer ,但不会受到 defer对变量的影响
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test1", err)
		}
	}()
	testsub()
}

/*
test1 : 0

*/

func main() {
	test()
}
