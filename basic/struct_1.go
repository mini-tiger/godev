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
	var st1 = &S1{}    //本地别名之后， 别名类型 只有别名自己 绑定的方法
	st1.A= map[int]string{
		1:"a",
	}

	var st2 = &utils.Struect_1{} //原始的 有原始绑定的方法
	fmt.Println(st2)
	fmt.Println(st2.Bind_str1_b())

}
