package main

import "fmt"

func Producer(ch chan<- int) {

	for i:=0;i<10;i++{
		ch <-i
		fmt.Printf("Producer %d\n",i)
	}
	close(ch)
}

func Consumer(ch <-chan int) {
	for i:=range ch{
		fmt.Printf("Consumer %d\n",i)
	}
}

func main() {

	var ch chan int = make(chan int, 5)
	go Producer(ch)
	Consumer(ch)
}
