package main

import (
	"fmt"
	"study/utils"
)

type t_struct struct {
	server *utils.Struct_1
}

func Newserver() (tStruct *t_struct) {
	ss := utils.Struct_1{map[int]string{1: "a"}, "b"}
	s := t_struct{server: &ss}

	return &s

}

func (self *t_struct) addflags(a string) {
	(*self).server.B = a
}

func main() {

	op := Newserver()
	fmt.Println(op.server.B)
	op.addflags("c")
	fmt.Println(op.server.B)

}
