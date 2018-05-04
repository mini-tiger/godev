package main

//!+
import "sync"
import (
	"fmt"
	// "time"
)

func main() {

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan struct{})
	for i := 0; i < 10; i++ { //10次只执行一次，初始化用
		go func() {
			once.Do(onceBody) //Do方法当且仅当第一次被调用时才执行函数
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

}
