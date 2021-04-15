package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"
)

var count int64 = 0
var maxGoroutines chan struct{} = make(chan struct{}, runtime.NumCPU())

func main() {
	for i := 'A'; i <= 'Z'; i += 2 {
		// 4, chan2 取出元素,执行打印两个字符 ,

		fmt.Print(string(i))
		fmt.Print(string(i + 1))
		// 5, chan1 接收一个元素,进入阻塞状态,等待取走元素,进入第2步,2345步一直循环直到打印完

	}
	st := time.Now()
	search("/home/go/")

	fmt.Println(time.Since(st))
	fmt.Println(count)

}

func search(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			search(filepath.Join(path, file.Name()))
		} else {
			atomic.AddInt64(&count, 1)
		}

	}

}
