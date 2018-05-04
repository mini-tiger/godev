package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func test(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	// waitgroup.Add(-1)
}

func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go func(i int) {
			fmt.Println(i)
			waitgroup.Done()
		}(i)
		// go test(i)
	}
	waitgroup.Wait() //.Wait()这里会阻塞主函数，直到队列中所有的任务结束就会解除阻塞
	fmt.Println("done!")
}
