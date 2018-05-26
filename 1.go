package main

import (
	"fmt"
	"sync"
)

var cc1 chan int = make(chan int)
var waitgroup sync.WaitGroup

var cc2 chan int = make(chan int)

func main() {
	f1()
	f11()
	f2()

	waitgroup.Wait()
}

func f1() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go func(i int) {
			cc1 <- i

		}(i)

	}
}

func f11() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go func(i int) {
			cc2 <- i

		}(i)

	}
}
func f2() {

	for {

		select {
		case e1 := <-cc1:
			fmt.Println(e1)
			fmt.Println("cc1")
			waitgroup.Done()
		case e2 := <-cc2:
			fmt.Println(e2)
			fmt.Println("cc2")
			waitgroup.Done()
		default:
			return

		}
	}
}
