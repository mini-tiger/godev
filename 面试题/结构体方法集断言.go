package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  结构体方法集断言
 * @Version: 1.0.0
 * @Date: 2021/5/18 上午10:07
 */

type Myint int

func (m Myint) addNoPointer(n int) { //struct 值传递， 不使用指针，会创建新对象
	m = m + Myint(n)
}

func (m *Myint) addUsePointer(n int) {
	*m = Myint(n) + *m
}

func main() {
	var a Myint = 1

	var b Myint = 1
	a.addNoPointer(1)
	b.addUsePointer(1) // 使用指针 改变自身
	fmt.Println(a)     // 1
	fmt.Println(b)     // 2

}
