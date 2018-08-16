package main

import (
	"fmt"
)

type i int
type s string

type PP struct {
	X i
	Y s
}

func main() {
	var (
		p1 PP
		p2 PP
	)

	var a = func(px1 *PP, px2 *PP) {
		fmt.Printf("px1: %T,px2: %d \n", *px1, *px2)
		fmt.Printf("px1 address: %p, px2 address %p \n", px1, px2)
		fmt.Println(*px1 == *px2)
		(*px1).Y = "2"
		fmt.Println(*px1 == *px2)
	}
	a(&p1, &p2)
	fmt.Println(p1 == p2)

}
