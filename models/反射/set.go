package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x) // 指针&x 的副本 不能寻址
	d := c.Elem()            // 指针提取，可以寻址
	fmt.Println(a.CanAddr(), a.CanSet())
	fmt.Println(b.CanAddr(), b.CanSet())
	fmt.Println(c.CanAddr(), c.CanSet())
	fmt.Printf("%v,%t,%t\n", d, d.CanAddr(), d.CanSet())
	/*
		false false
	false false
	false false
	2,true true
	*/
	fmt.Println(x)
	d.Set(reflect.ValueOf(4)) //设置
	fmt.Println(x)

	set_ex()
}
func set_ex() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	/*
	func (v Value) Elem() Value
		Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。

	*/
	rx.SetInt(22) // setint,注意类型rx.SetString("111")//错误

	fmt.Println(x)
	rx.Set(reflect.ValueOf(33)) // setint 作用相同
	fmt.Println(x)

	xx := 2
	set_interface(xx)

}

func set_interface(xx interface{}) {
	fmt.Println(xx)
	rxx := reflect.ValueOf(&xx).Elem()
	//rxx.SetInt(4)//xx 是接口，则不能使用setint,setstring等方式
	rxx.Set(reflect.ValueOf(55))
	fmt.Println(xx)

}
