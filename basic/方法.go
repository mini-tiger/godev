package main

import (
	"fmt"
)

type S1 struct {
	X, Y int
}

func (s1 S1) ss1(x, y int) int {
	return s1.X + s1.Y + x + y

}

type S2 struct {
	X, Y int
}

func (s2 *S2) ss2(x, y int) int {
	return s2.X + s2.Y + x + y
}

type A1 []int

func (a1 A1) aa1(a2 int) int {
	return a1[0] + a2

}

func main() {
	//绑定方法

	var a A1
	a = []int{1, 2}
	fmt.Println(a.aa1(2))

	fmt.Println(S1{1, 2}.ss1(1, 2))

	s := &S1{1, 2}
	fmt.Println(s.ss1(1, 2)) //接收为 类型， 自动取 内存地址的值

	//绑定方法 指针
	fmt.Println((&S2{1, 2}).ss2(1, 2))

	ss := S2{1, 2}
	fmt.Println((&ss).ss2(1, 2))

	fmt.Println(ss.ss2(1, 2)) //接收是指针  自动 取内存地址
	//嵌套

	sss := S3{&ss, 2}
	fmt.Println((&sss).ss3()) //sss.ss3()

	sss.X = 5                  //sss.S2.X=5
	fmt.Println(sss.ss2(1, 2)) //继承 子S2 绑定方法

	sss1 := S3{&S2{1, 2}, 3}
	fmt.Println(sss1.S2.ss2(1, 2)) // sss1.ss2(1,2)
}

type S3 struct {
	*S2
	Z int
}

func (s3 *S3) ss3() int {
	return s3.X + s3.Y + s3.Z

}
