package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	runtime.GOMAXPROCS(0)

	var c chan interface{}
	fmt.Println(unsafe.Sizeof(c))
}
