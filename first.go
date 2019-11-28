package main

import (
	"fmt"
	"strings"
)

func main()  {
	s:="dd d \n   ddd "
	ss:=&s
	fmt.Println(*ss)

	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(*ss,"\n","")," ",""))
}