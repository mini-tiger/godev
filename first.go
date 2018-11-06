package main

import "fmt"

type ff struct {
	ffunc []func()string
	i int
}

func a() string {
	return "a"
}

func b()  string {
	return "b"
}

func main()  {
	f:=ff{[]func()string{a,b},1}
	for _,i:=range f.ffunc{
		fmt.Println(i())
	}
}