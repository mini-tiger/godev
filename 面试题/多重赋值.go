package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  多重赋值
 * @Version: 1.0.0
 * @Date: 2021/5/18 上午9:48
 */

func main() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
	/*
	   多重赋值分为两个步骤，有先后顺序：
	   计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
	   赋值；
	   xxx 所以本例，会先计算左边 s[i-1]， 取i s[0]地址，之后等号右边是两个表达式是常量，所以赋值运算等同于 i, s[0] = 2,
	   "Z" 。
	*/

}
