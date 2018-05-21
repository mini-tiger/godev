package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	for _, v := range s {
		go func() { // 所有匿名函数都是传入 外层变量的地址，执行时最后v是c,所以输出3次c
			fmt.Println(v)
		}()
	}

	for _, s1 := range s {
		go func(s string) { //当作参数传入
			fmt.Println("variables input", s)
		}(s1)
	}

	for _, ss := range s { //单线程
		fmt.Println("singnal thread", ss)
	}

	select {} //无限等待
}

/*
[a b c]
c
c
c
*/
