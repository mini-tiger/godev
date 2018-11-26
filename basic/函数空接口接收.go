package main

import (
	"fmt"
	"reflect"
	"strings"
)

type a1 struct {
	a int
}
type b2 struct {
	b int
}

func abc(args interface{})  {
	s:=reflect.TypeOf(args)
	fmt.Println("============",s.String())
	fmt.Println("============",s.Name())
	//ss:=args.(a)
	switch {
	case strings.Contains(s.String(),"a1"):
		ss:=args.(*a1)

		fmt.Println(ss.a)
	case strings.Contains(s.String(),"b2"):
		ss:=args.(*b2)

		fmt.Println(ss.b)
	}

	//fmt.Print(ss)
}
func main()  {
	var aa a1
	aa.a=1

	var bb b2
	bb.b=2


	abc(&aa)
	abc(&bb)
}

