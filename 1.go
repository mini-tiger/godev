package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker_adder(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
	sw.Done()
	//fmt.Println("end")
}

var sw sync.WaitGroup

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	sw.Add(2)
	// spin up a goroutine.
	go worker_adder(s[:len(s)/2], c1)
	// spin up a goroutine.
	go worker_adder(s[len(s)/2:], c2)

	x, y := <-c1, <-c2 // receive from c1 aND C2
	//x, _:= <-c1
	// 输出从channel获取到的值
	fmt.Println(x, y)

	fmt.Println(runtime.NumGoroutine())
	sw.Wait()
	fmt.Println(runtime.NumGoroutine())
}
