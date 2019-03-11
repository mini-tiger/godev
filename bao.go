package main

import (
	"time"
	"fmt"
)



type I interface {
	T()int
}

type ii struct {
	II int
}

func (ii *ii)T() int {
	return ii.II
}

var c1 chan int=make(chan int,0)
var c2 chan I=make(chan I,0)


func main()  {
	go func(c chan<- int) {
		for i:=0;i<10;i++{
			c<-i
			time.Sleep(1*time.Second)
		}
	}(c1)
	go func(c <-chan int,cc chan<- I) {
		for{
			select {
			case i:=<-c1:
				var i1 I
				i1=&ii{i}
				cc<-i1
			}
		}
	}(c1,c2)
	go func() {
		for{
			select {
			case i:=<-c2:
				fmt.Println(i.T())
			}
		}
	}()
	select {}
}