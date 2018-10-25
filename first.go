package main

import (
	"fmt"
	"time"
)

var c1 = make(chan int,4)

func main()  {
	go func() {
		for  {
			if dom,ok:=<-c1;ok{
				fmt.Println(dom)
			}else {
				break
			}
		}
	}()

	go func() {
		for{
			c1<-1
			time.Sleep(time.Duration(1)*time.Second)
		}
	}()
	select {}
}