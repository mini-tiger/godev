package main

import (
	"fmt"
	"unicode/utf8"
)

var a string = "我是 taojun"

func main() {
	var ab []byte = []byte(a) //byte 是int8的别名，基本上一个位置代表一个ascii码值，所以中文一个要占3个位置
	fmt.Printf("string length:%d,value:%v\n", len(a), a)
	fmt.Printf("[]byte length:%d,value:%v\n", len(ab), ab)
	for i, v := range a { //中文一个字，占三个字节
		fmt.Printf("index:%d,value:%c\n", i, v) //%c 该值对应的unicode码值
	}

	var ar []rune =[]rune(a) //rune 是int32的别名，一个位置能代表所有单个字符,所以中文占一个位置
	fmt.Printf("[]rune length:%d,value:%c\n", len(ar), ar)
	for ii,vv:=range ar{
		fmt.Printf("index:%d,value:%c\n", ii, vv) //%c 该值对应的unicode码值
	}
	fmt.Println(utf8.RuneCountInString(a)) //不能转换[]rune，计算字符串 长度
}
