package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  原始结构
 * @Version: 1.0.0
 * @Date: 2021/6/9 上午10:14
 */

func main() {
	var s []int = []int{1, 2, 3}
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%+v\n", ss)

	var str string = "123"
	strs := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("%+v\n", strs)

	// []byte string 转换
	fmt.Println("!!!! []byte string 转换")
	var x []byte = []byte("AB")
	fmt.Println(string(x))
	fmt.Println(*((*string)(unsafe.Pointer(&x)))) // unsafe.Pointer 全能型指针

	fmt.Println("!!!! string  []byte 转换")

	var xx string = "ABC"
	fmt.Println([]byte(xx))
	fmt.Println(*((*[]byte)(unsafe.Pointer(&xx))))

	fmt.Println("!!!!   struct []byte 转换")
	type A struct {
		AA int
	}
	var AAA = &A{1}
	var sizeof int = int(unsafe.Sizeof(*AAA)) // 测量struct size

	var xxx reflect.SliceHeader
	xxx.Cap = sizeof
	xxx.Len = sizeof
	xxx.Data = uintptr(unsafe.Pointer(AAA))
	xxxbyte := *(*[]byte)(unsafe.Pointer(&xxx))
	fmt.Println(xxxbyte)

	fmt.Println("!!!!  []byte  struct 转换")
	xxxbyteToStruct := (*A)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&xxxbyte)).Data))
	fmt.Println(xxxbyteToStruct)

}
