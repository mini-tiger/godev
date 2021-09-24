package main

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
	d := B{}
	d.AA
}
