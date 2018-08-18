package main

import (
	"fmt"
	"godev/g"
)

func main()  {
	fmt.Println(g.Age,g.Name)
	var s1 *g.Struct_1 = &g.Struct_1{1}
	fmt.Printf("s1: %+v",*s1)
}
