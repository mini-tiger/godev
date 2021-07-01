package main

import "fmt"

func main() {
	a := 1
	b := []int{1, 2, 3}
	a, b[a-1] = 2, 2
	fmt.Println(a)
	fmt.Println(b)
	aa := make(map[int]int, 0)
	aa[1] = 1
	delete(aa, 1)
	aaxx()
}

func aaxx() (x interface{}) {
	fmt.Println(x == nil)
	return
}
