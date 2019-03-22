package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Println(a)
	//前置补0
	fmt.Printf("%02d", a)
	fmt.Println("")
	fmt.Printf("%0*d", 3, a)
}