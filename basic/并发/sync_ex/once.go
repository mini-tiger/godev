package main

//!+
import "sync"
import (
	"fmt"
	// "time"
	"time"
)
var done = make(chan struct{},0)

func main() {

	var once *sync.Once = new(sync.Once)
	onceBody := func() {
		fmt.Println("Only once")
		done <- struct{}{}
	}

	for i := 0; i < 10; i++ { //循环10次，但do方法只会在第一次调用时执行
		go func() {
			once.Do(onceBody) //Do方法当且仅当第一次被调用时才执行函数
		}()
		time.Sleep(time.Duration(1)*time.Second)
	}

	for i := 0; i < 10; i++ {
		<-done
		fmt.Printf("第 %d 次<-done\n",i)
		time.Sleep(time.Duration(1)*time.Second)
	}
	select {}
}
