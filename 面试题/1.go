package main

import (
	"fmt"
)

func main() {
	var c chan int = make(chan int, 1)
	c <- 1
	close(c)
	fmt.Println(<-c)
}
