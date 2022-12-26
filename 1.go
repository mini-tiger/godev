package main

import "fmt"

type Ai interface {
	ff(string) string
}

type Aii interface {
	ff(string) string
	Abc() int
}

type T int

func (t *T) ff(s string) string {
	return s
}
func (t *T) Abc() int {
	return int(*t)
}

func main() {

	t := T(1)
	fmt.Println(t.Abc())
	a := Aii(&t)
	fmt.Println(a.Abc())

	aa := a.(Ai)
	fmt.Println(aa.ff("11"))
}
