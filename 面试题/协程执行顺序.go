package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d ", i)
		}(i)
	}

	go func() {
		defer wg.Done()
		fmt.Println("这里这里是最后一个准备的协协，先执行")
	}()

	wg.Wait()
	// eee 输出0 1 2 3
	// xxx 最先执行最后中一个好的协程,，其它按顺序执行
}
