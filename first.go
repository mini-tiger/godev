package main

import (
	"fmt"
	"time"
)

var tc chan int = make(chan int, 100)

func send() {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("begin tc", time.Now())
		tc <- 1
		fmt.Println("end tc", time.Now())
	}
	close(tc)

}
func recv() {
	for {
		select {
		case i, ok := <-tc:
			fmt.Println(i, ok, time.Now())
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go recv()
	go send()

	select {}
}
