package main

import (
	"fmt"
	"reflect"
)

type s struct {
	A int
}

func main()  {
	var ss s
	ss.A=1
	get(&ss)
}

func get(i interface{})  {
	typ:=reflect.TypeOf(i)
	v:=reflect.ValueOf(i)
	fmt.Println(typ.Kind() == reflect.Ptr,typ.Name())
	fmt.Println(v.Type(),v.Kind() == reflect.Ptr)
	vv:=v.Type()
	fmt.Println(vv.Kind(),vv.Name(),vv.PkgPath(),vv.String())
	//fmt.Println(v.in)
	ss:=i.(*s)
	fmt.Printf("%T\n",ss)
}