package main

import "fmt"

func main() {
	aa := []int{1, 2, 3}
	fmt.Println(aa)
	test(aa...)

}

func test(a ...int) {
	fmt.Println(a)
}
