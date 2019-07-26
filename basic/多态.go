package main

import (
	"fmt"
	"strconv"
)

type DT interface {
	Run(string)
}

type A struct {
	aa int
}

func (a *A)Run(s string) {
	fmt.Println(s)
}


func Ex(a []DT)  {
	for i,v:=range a{
		//fmt.Printf("%d,%s",i,)
		v.Run(strconv.Itoa(i))
	}
}


func main()  {
	//a.run("abc")
	a:=&A{1}
	Ex([]DT{a})

}
