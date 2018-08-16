package main

import (
	"fmt"
	config "study/g"
	"study/utils"
	//"github.com/spf13/pflag"
)

// type i int
// type s string

// type PP struct {
// 	X i
// 	Y s
// }

type S1 utils.Struct_1

func (self *S1) bind_Str1() string {
	fmt.Println(self.A)
	return "a"
}
func main() {

	defer fmt.Println("defer end")
	fmt.Println(config.Ip, config.DB)
	var st1 = &S1{} //本地别名之后， 别名类型 只有别名自己 绑定的方法
	st1.A = map[int]string{
		1: "a",
	}

	var st2 = &utils.Struct_1{}
	fmt.Println(st2)
	fmt.Println(st2.Bind_str1_b())
	ex()
}

type st1 struct {
	a, b int
	c    []int
}

func ex() {
	var st11 *st1 = &st1{}
	st11.a = 1
	st11.b = 1
	st11.c = []int{1}
	fmt.Println(*st11)
	//
	st22 := &struct {
		a, b int
		c    []int
	}{1, 1, []int{1, 2}}
	fmt.Println(*st22)

}
