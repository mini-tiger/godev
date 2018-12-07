package main

import (
	"time"
	"fmt"
)

var tt *time.Ticker
var tr *time.Timer // 一次性
func main() {
	tt = time.NewTicker(time.Duration(2) * time.Second)
	tr = time.NewTimer(time.Duration(5) * time.Second)
	//go func() {
	//	for {
	//		<-tt.C
	//		fmt.Println(time.Now())
	//	}
	//}()
	//go func() {
	//	<-tr.C
	//	fmt.Println("this is Timer")
	//}()
	//}
	for {
		select {
		case <-tt.C:
			fmt.Println(time.Now())
		case <-tr.C:
			fmt.Println("this is Timer")
		}
	}
	select {}
}