package main

import (
	"time"
	"fmt"
	"log"
)

var tt *time.Ticker
var tr *time.Timer // 一次性
func main() {
	fmt.Println(time.Now())
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
	for i := 0; ; i++ {
		select {
		case t1 := <-tt.C:
			log.Println("Next Time", t1.Add(time.Duration(2)*time.Second))
			if i > 3 {
				tt.Stop() //  todo 可以停
			}
		case t2 := <-tr.C:
			fmt.Println("this is Timer", t2)
		}
	}

	select {}
}
