package main

import (
	"fmt"
	"strings"
)

type A struct {
	S *strings.Builder
}
func main()  {
	a:=A{}
	a.S=&strings.Builder{}
	a.S.Grow(10)
	a.S.WriteString("1")
	fmt.Println(a.S.String())
}