package main

import "fmt"

type A1 struct {
	a []map[string]interface{}
	b func(int)
}

func main()  {

	a1 := A1{}
	a1.a= make([]map[string]interface{},0)
	a1.b= func(a int) {
		fmt.Println(a)
	}

	fmt.Println(a1)
	a1.b(22)//
}