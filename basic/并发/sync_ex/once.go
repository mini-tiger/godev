package main

import (
	"fmt"
	"sync"
	"time"
)

var done = make(chan struct{}, 0)

func main() {

	//for i := 0; i < 10; i++ {
	//	<-done
	//	fmt.Printf("第 %d 次<-done\n",i)
	//	time.Sleep(time.Duration(1)*time.Second)
	//}
	//select {}
	for i := 0; i < 10; i++ { //循环10次，但do方法只会在第一次调用时执行
		tt()
		time.Sleep(time.Duration(1) * time.Second)
	}

}
func tt() {

	var once *sync.Once = new(sync.Once)
	onceBody := func() {
		fmt.Println("Only once")
		done <- struct{}{}
	}

	go func() {
		once.Do(onceBody) //Do方法当且仅当第一次被调用时才执行函数
	}()

}
