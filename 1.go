package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var s string

func main() {
	s = "!23"
	fmt.Printf("%v,%p\n", s, &s)
	//sl:=strToSlice(&s)
	sl := strings.Split(s, "")
	fmt.Printf("%v,%p\n", sl, &sl)

}
func strToSlice(s *string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(s))
	//fmt.Printf("%v,%p\n",sl,sl)
	ss := reflect.SliceHeader{
		Len:  str.Len,
		Data: str.Data,
		Cap:  str.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&ss))
}
