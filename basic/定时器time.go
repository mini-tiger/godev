package main

import (
	"time"
	"fmt"
)

func method_1(chan1 *time.Ticker)  {
	for {
		select {
		case out,ok := <-chan1.C:
			if ok {
				fmt.Printf("this is method1 %s \n",out)
			}
		}
	}
}

func method_2()  {
	d := time.Duration(1 * time.Second)
	tick := time.Tick(d)
	for {
		select {
		case out:=<-tick:
			fmt.Printf("this is method2 %s \n",out)
		}
	}
}

func method_3(chan2 *time.Timer)  {
	for {
		select {
		case out,ok := <-chan2.C:
			if ok {
				fmt.Printf("this is method3 one task %s \n",out)
			}
			break
		}
	}
}

func main()  {
	chan1 := time.NewTicker(1*time.Second)
	go method_1(chan1)
	go method_2()

	//闹钟 只一次
	chan2 := time.NewTimer(3*time.Second)
	go method_3(chan2)
	select {}
}
