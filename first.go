package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main()  {
	var i int32
	var i32 uint16
	fmt.Printf("%v,%d\n",&i,i)
	atomic.AddInt32(&i,1)
	fmt.Printf("%v,%d\n",&i,i)

	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(i32))
}