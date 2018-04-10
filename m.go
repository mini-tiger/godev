package main

import (
	b "bao_demo"
	b1 "bao_demo/bao1"
	. "fmt"
)

func main() {
	var (
		x int = 1
	)
	Println("2")
	Println(b.Add(new(int))) //共享函数 首字母大写
	Println(b.T1(x))
	Println(b.AA)
	Println(b1.CC)
}
