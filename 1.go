package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type All interface {
	Get() int
	Set()
}
type G interface {
	Get() int
}

type S int

func (s *S) Get() int {
	return 1
}
func (s *S) Set() {
	return
}

func byte2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func string2byte(b string) []byte {
	return *(*[]byte)(unsafe.Pointer(&b))
}
func main() {
	runtime.GOMAXPROCS(1)
	var b []byte
	b = []byte("Hello")
	fmt.Println(byte2string(b))

	fmt.Println(string(string2byte(byte2string(b))))
}
