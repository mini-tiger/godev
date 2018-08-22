package main

import (
	"godev/mymodels/set"
	"fmt"
)

func main()  {
	ss:=set.Newset()
	//fmt.Println(ss)
	ss.Add("aa")
	ss.Add("bb")
	ss.Add("aa")
	fmt.Println(ss.Exists("cc"))
	fmt.Println(ss.Keys())
}
