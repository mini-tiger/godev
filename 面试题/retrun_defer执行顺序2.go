package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  retrun_defer执行顺序
 * @Version: 1.0.0
 * @Date: 2021/4/28 上午10:03
 */

func trace(s string) string {
	fmt.Println("trace :", s)
	return s
}

func un(s string) string {
	fmt.Println("un :", s)
	return s
}

func main() {
	fn1()

	fn2()
}

/*
结果:
in fn1
trace : a
un : a
*/

func fn1() { //
	defer un(trace("a")) // 提前准备好 ，先执行 trace("a")

	fmt.Println("in fn1")

}

func fn2() { //
	defer func() {
		un(trace("a")) // 提前准备好， 函数里面不用执行
	}()

	fmt.Println("===========in fn2")

}
