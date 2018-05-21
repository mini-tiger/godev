package main

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...create abort channel...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	select { //如果有多个case为真，随机执行一个
	case <-time.After(5 * time.Second): //After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C。
		fmt.Println("已经等待5秒")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
	select {} //无限等待

}

//!-

func launch() {
	fmt.Println("Lift off!")
}
