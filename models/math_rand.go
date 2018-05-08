package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Float64())
	fmt.Println(rand.Int())

	//使用给定的种子创建一个伪随机资源。
	a := rand.NewSource(111111)
	b := rand.New(a)
	fmt.Println(b.Int())
}
