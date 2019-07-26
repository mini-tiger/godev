package main

import "fmt"

type AI int

var C1 chan AI =make(chan AI ,0)

func send()  {
	for i:=0;;i++{

		C1<-AI(1)

	}
}

func recv()  {
	for{
		select {
		case i:=<-C1:
			fmt.Printf("%T,%v\n",i,i)
		}

	}
}

func main()  {
	go recv()
	go send()
	select {}

}
