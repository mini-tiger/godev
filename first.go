package main

import (
	"sync"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

var sy *sync.WaitGroup = new(sync.WaitGroup)

func main()  {
	//fmt.Println(sy)
	sy.Add(1)
	go func() {
		for {
			sy.Add(1)
			//fmt.Println("this is 1")
			//time.Sleep(time.Duration(1)*time.Second)
			}
	}()

	go func() {
		for  {
			fmt.Println("this is 2")
			sy.Done()
			time.Sleep(time.Duration(1)*time.Second)
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()
	sy.Wait()
}