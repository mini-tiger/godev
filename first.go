package main

import (
	"fmt"
	"sync"
	"time"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func main() {
	var l sync.WaitGroup
	go func() {
		l.Add(1)
		<-balances
		fmt.Println(1)
		l.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(2)
		balances <- 11
	}()
	time.Sleep(3 * time.Second)
	l.Wait()

	c := make(chan bool)
	go func() {
		// fmt.Println(1)
		c <- true
	}()

	<-c
}
