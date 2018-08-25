package main

import "fmt"

type aa struct {
	A int
}

func main() {
	a := new(aa)
	a.A = 1
	b := new(aa)
	b.A = 2
	fmt.Println("%p,%T", &a, b)

	a = new(aa)
	a.A = 2
	b = new(aa)
	b.A = 3
	fmt.Println("%T,%T", a, b)

}
