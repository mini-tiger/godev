package main

import (
	"fmt"
	"study/utils"
)

type Iface_1 interface {
	look() string
}

type Struct_2 struct {
	B string
}

func (self *Struct_2) look() string {
	return "this is Struct_2 look func"
}

func (self *Struct_2) close() string {
	return "this is Struct_2 close func"
}

func main()  {
	fmt.Println()

	var s2 Struct_2
	s2.B="B"
	fmt.Println(s2.look())
	fmt.Println(s2.close())

	var iface1 Iface_1
	iface1 = &s2//只要有 look方法都可以 挂在 接口上
	fmt.Println(iface1.look()) //接口只有 look方法，没有其它方法，结构体和接口要在同一个文件

	//utils.Inter() // 调用其它文件的接口方法
	}