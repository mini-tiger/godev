package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  结构体
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午11:26
 */

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {
}
func main() {
	var s *Student
	if s == nil { // 只有类型 为nil,指针类型， nil 指针
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	/*
		当且仅当动态值和动态类型都为 nil 时，接⼝类型值才为 nil。上⾯的代码，给变量 p 赋值之后，p 的动态值是
		nil，但是动态类型却是 *Student，是⼀个 nil 指针，所以相等条件不成⽴。
	*/

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
