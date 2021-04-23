package main

import (
	"fmt"
	"os"
	"strconv"
)

var chstr1 chan string=make(chan string,0)
var chstr2 chan string=make(chan string,0)
var exitchan chan struct{}=make(chan struct{},0)

func main()  {
	str1:=[]string{"A","B","C","D","E","F","G"}
	str2:=make([]string,6)
	for i:=0;i<6;i++{
		str2[i]=strconv.Itoa(i)
	}
	//fmt.Println(str1)
	//fmt.Println(str2)


	 go func() {
	 	for {
			select {
	 		case a,_:=<-chstr1:
	 			fmt.Println(a)
			case a,_:=<-chstr2:
				fmt.Println(a)
			}
		}
	 }()
	go func() {
		<-exitchan
		os.Exit(0)
	}()
	for i,value:=range str2{
		chstr1<-value
		chstr2<-str1[i]
	}
	exitchan<- struct{}{}


}