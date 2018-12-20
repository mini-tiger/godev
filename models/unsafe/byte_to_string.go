package main

import (
	"fmt"
	"unsafe"
	"strings"
)

func main()  {
	b:=[]byte("this is byte")
	//fmt.Println(string(b)) // 此方法 会复制[]byte
	s := *(*string)(unsafe.Pointer(&b))
	s = "  baad  "
	s = strings.Trim(s," ")
	fmt.Println(s) // 推荐方法，[]byte 改变，而随着变

}
