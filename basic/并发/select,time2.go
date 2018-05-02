package main

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...create abort channel...

	//!-

	abort := make(chan struct{}) //类型无所谓，struct{} 不占内存
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second) //定时器,间隔发送到通道
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick: //通过上面定时器 每秒 可以取出
			// Do nothing.
		case <-abort: //如果 按了键盘上的键回车
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
