package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3}
	fmt.Printf("sl1 address %p\n", sl1)

	a := 11
	var i *int
	i = &a
	fmt.Println(*i)
}
