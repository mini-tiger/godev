package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
}

func a(i ...int) {
	fmt.Println(i)
}
