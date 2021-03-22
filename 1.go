package main

import "fmt"

func p() {
	i := 1
	p := &i
	defer fmt.Println((*p) * 2)
	i = 2
	fmt.Println(*p)
	panic(1)
}

func main() {
	p()
}
