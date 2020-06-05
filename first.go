package main

import (
	"fmt"
	"strings"
)

type A struct {
	S strings.Builder
}
func main()  {
	var tmp strings.Builder
	tmp.Grow(10)
	tmp.WriteString("1")
	fmt.Println(tmp.String())
}