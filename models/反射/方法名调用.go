package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type A struct {
	A1 int

}

func (a *A)Get() int {
	return a.A1
}

func (a *A)Get1(aa int) string {
	return strconv.Itoa(a.A1 + aa)
}

func main() {
	aa:=new(A)

	WriteLog(aa)

	fmt.Println(callReflect(aa,"Get1",1)[0].String())
}

func WriteLog(o interface{}) {
	aa:=make([]reflect.Value,0)
	v := reflect.ValueOf(o)
	f := v.MethodByName("Get")
	fmt.Println(f.Call(aa)[0])
}

func callReflect(any interface{}, name string, args... interface{}) []reflect.Value{
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	}else {
		return v.Call(inputs)
	}

}