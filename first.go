package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  first
 * @Version: 1.0.0
 * @Date: 2021/9/10 下午2:56
 */

type A struct {
	AA int
}
type B struct {
	A
}

func (a *A) sss() int {
	return 1
}

func main() {
	a := A{1}
	a1 := a
	a1.AA = 2
	fmt.Println(a, a1)

	a2 := &a
	a2.AA = 3
	fmt.Println(a, a1, a2)

}
