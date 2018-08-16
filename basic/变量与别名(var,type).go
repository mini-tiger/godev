package main

import "fmt"

var i int
type alias_i int

func main()  {
	i  = 1
	fmt.Println(i)

	ii := new(int)
	*ii = i
	fmt.Println(*ii)

	ai := alias_i(1)
	iii := new(alias_i)
	*iii = ai
	fmt.Println(*iii)

	iiii := new(alias_i)
	*iiii = i    // 这里不会执行，别名以后 ，不能用之前的类型赋值
	fmt.Println(*iiii)
}