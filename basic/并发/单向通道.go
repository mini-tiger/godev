package main

import (
	"fmt"
	// "time"
)

//!+
func counter(put chan<- int) {
	for x := 0; x < 10000; x++ {
		put <- x
	}
	// time.Sleep(1 * time.Second)
	defer close(put)
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

func t() {
	naturals := make(chan int, 10)
	squares := make(chan int, 10)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
func main() {
	t()
	// time.Sleep(11 * time.Second)
}
