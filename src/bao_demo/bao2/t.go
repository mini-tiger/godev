package bao2

import (
	"fmt"
)

func AA(a *int) {
	fmt.Println(*a)
	fmt.Println("this is bao _import")
}

func init() {
	const (
		a = 1
		B = 2
		S = "11"
	)
	aa := 1
	AA(&aa)
}
