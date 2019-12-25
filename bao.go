package main

import (
	"time"
)

var tmpLock = make(chan struct{}, 0)

func main() {

	go func() {
		for {
			aa()
			time.Sleep(time.Duration(2 * time.Second))
		}
	}()
	go func() {
		for {
			<-tmpLock
			time.Sleep(time.Duration(1 * time.Second))
		}
	}()

	select {}
}

func aa()  {
	defer func() {
		tmpLock <- struct{}{}
	}()
}
