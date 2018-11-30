package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	b:=[]byte("this is byte")
	//fmt.Println(string(b)) // 此方法 会复制[]byte
	fmt.Println(*(*string)(unsafe.Pointer(&b))) // 推荐方法，[]byte 改变，而随着变

}
