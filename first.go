package main

import "fmt"

var c chan int

func main() {

	c <- 1

	select {
	case <-c:
		fmt.Println(1)
	}
}
