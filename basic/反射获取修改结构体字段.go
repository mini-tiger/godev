package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	Name   string  `json:"name"`
	Age       int	`json:"age"`
}

func SetValueToStruct() *Person {
	p := &Person{Age:11}
	v := reflect.ValueOf(p).Elem()
	vn:=v.FieldByName("Name")
	fmt.Printf("%T\n",vn.String())
	fmt.Println(vn.String() == "")
	fmt.Println(v.FieldByName("Age"))

	return p
}


func main()  {
	p := SetValueToStruct()
	fmt.Println(*p)
}
