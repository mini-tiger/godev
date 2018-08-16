package utils

import (
	"fmt"
	"runtime"
)

var Abort chan int = make(chan int)
var CanLunch chan int= make(chan int)

func init() {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Printf("this is filename: %s \n", filename)
}
 // struect
type Struct_1 struct {
	A map[int]string
	B string
}

func (self *Struct_1) Bind_str1_b() string  {
	return "this is Bind_str_b"
}


//interface



type Iface_1 interface {
	look() string
}

type Struct_2 struct {
	B string
}

func (self *Struct_2) look() string {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Printf("this is filename: %s \n", filename)
	return "this is utils.Struct_2 look func"
}

func Inter()  {
	var U Iface_1
	var S2 Struct_2
	S2.B="BB"
	U = &S2
	fmt.Println(U.look())
}

