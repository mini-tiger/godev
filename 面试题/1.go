package main

import (
	"fmt"

	"unsafe"
)

type A struct {
	AA int
	BB int
}

func main() {
	var a *A = &A{1, 1}
	fmt.Println(a)
	b := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(unsafe.Offsetof(a.AA))))
	*b = 2
	fmt.Println(a)

}
