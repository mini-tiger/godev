package main

import "fmt"

func add(a,b int) int {
	return a+b
}

type  N struct {
	A func(int,int) int
}

func main()  {
	aa:=1
	bb:=2
	fmt.Println("=",add(aa,bb))
	fmt.Printf("add pointer %p\n",add)
	n:=N{add}
	fmt.Println(n.A(aa,bb))
}