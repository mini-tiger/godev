package main

import "fmt"

type fff struct {
	f1 []func(int) int
	interval int
}

var F []fff

func test1(a int) int {
	return a
}
func test2(aa int) int  {
	return aa
}

func initFunc()  {
	F=[]fff{
		{
			f1: []func(int) int{
				test1,
			},
			interval: 1,
		},
		{
			f1:[]func(int) int{
				test2,
			},
		},

	}
}
func runFunc()  {
	for _,v :=range F{
		for i,vv:=range v.f1 {
			fmt.Println(vv(i),v.interval)
		}
	}
}
func main()  {
	// 函数存入切片中的结构体
	initFunc()
	runFunc()
	fmt.Println("=================================================")
	// 函数存入切片
	initFS()
	runFS()

}
var FS []func(int) int
func initFS()  {
	FS=[]func(int)int{
		test1,
		test2,
	}
}
func runFS()  {
	for i,v := range FS{
		fmt.Println(v(i))
	}
}