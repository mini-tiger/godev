package main

import (
	"fmt"
	"time"
)

//!+
func counter(put chan<- int) {
	for x := 0; x < 10000; x++ {
		put <- x
	}
	time.Sleep(1 * time.Second)
	close(put)
}

func squarer(put chan<- int, pop <-chan int) {
	for v := range pop {
		put <- v * v
	}
	close(put)
}

func printer(pop <-chan int) {
	for v := range pop {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
