package main

import (
	"fmt"
	"os"
	// "sync"
	"text/tabwriter"
	"time"
)

type ss struct {
	c, r int
}

var lingpai = make(chan struct{}, 20)
var result = make(chan ss, 20)
var done = make(chan struct{})

var lingpai_plus = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	for i := 0; i < 20; i++ {
		lingpai <- struct{}{}

	}

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 等待键盘输入，则退出
		close(done)                    //close 后可以 取出空
	}()

	go func() { //令牌使用完，可以追加
		for {
			if cancelled() { //如果键盘输入为真，返回结束
				return
			}
			select {
			case <-lingpai_plus:
				lingpai <- struct{}{}
			}

		}

	}()

	tick := time.Tick(2 * time.Second)
	n := new(int)
	*n = 0

loop:
	for {

		select {
		case <-done: //键盘输入任意则停止
			break loop
		case <-lingpai:

			go suan(result)
		case <-tick:
			*(n)++
			fmt.Println(*n, len(result))
			for i := len(result); i > 0; i-- {
				printer(n, <-result)

			}

		}
	}

}

func printer(n *int, rr ss) {
	fmt.Fprintf(os.Stdout, "这是第%d次\n", *n)
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0) //https://studygolang.com/pkgdoc
	fmt.Fprintf(tw, format, "cishu", "result")
	fmt.Fprintf(tw, format, "-----", "------")
	// for t := range result {
	fmt.Fprintf(tw, format, rr.c, rr.r)
	// }
	tw.Flush() // calculate column widths and print table
}

func suan(rr chan<- ss) {

	var s ss
	for i := 0; i < 5; i++ {
		if cancelled() {
			return
		}
		time.Sleep(1 * time.Second)
		s.c++
		s.r += i * i
		// fmt.Println(i)
		result <- s
	}

	lingpai_plus <- struct{}{} //发送，此任务结束，令牌使用完，可以追加令牌
	return
}
