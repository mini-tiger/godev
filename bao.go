package main

import (
	"reflect"
	"unsafe"
)

func main()  {
	var a reflect.StringHeader
	a.Data=unsafe.Pointer("123")
}