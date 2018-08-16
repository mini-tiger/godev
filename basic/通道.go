package main

import (
	"fmt"
	"time"
	"study/utils"
)

func launch() {
	fmt.Println("nuclear launch detected")
}

func commencingCountDown(canLunch chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 20; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<- c
	}
	canLunch <- -1
}

func isAbort(abort chan int) {
	time.Sleep(2*time.Second)
	abort <- -1
}

func main() {
	fmt.Println("Commencing coutdown")


	
	go isAbort(utils.Abort)
	go commencingCountDown(utils.CanLunch)
	select {
	case <- utils.CanLunch:

	case <- utils.Abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}