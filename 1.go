package main

import "fmt"

type HH interface {
	Abc() int
	Bcd(func(int) int) int
}

type H interface {
	Abc() int
}

type A int

func (a *A) Abc() int {
	return int(*a)
}
func (a *A) Bcd(f func(int) int) int {
	return f(int(*a))
}

func main() {
	a := A(1)
	h := H(&a)
	h.Abc()
	fmt.Printf("%T\n", h)
	hh := h.(HH)
	fmt.Printf("%T\n", hh)

	fmt.Println(hh.Bcd(func(i int) int {
		return i
	}))

}
