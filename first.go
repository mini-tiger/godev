package main

import (
	"fmt"
)

var c1 chan int = make(chan int, 2)

func main() {
	fmt.Println(len(c1))
	A()
	c1 <- 10
	recv(c1)

	//for i:=0;i<10;i++{
	//
	//	if i%2==0{
	//		c1<-i
	//	}else{
	//		c1<-i
	//	}
	//}

}

func recv(c chan int) {
	for {
		fmt.Println(len(c1))
		select {
		case i := <-c:
			fmt.Println(i)
			fmt.Println(len(c1))
			fmt.Println("===========")
		}
	}

}
