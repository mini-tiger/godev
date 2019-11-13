package main

import (
	"log"
	"runtime/debug"
	"time"
)

func f() {
	container := make([]int, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}

func main() {
	log.Println("start.")
	f()

	log.Println("force gc.")
	time.Sleep(5 * time.Second) // 保持程序不退出
	runtime.GC() // 调用强制gc函数
	log.Println("force FreeOsMemory.")
	time.Sleep(5 * time.Second) // 保持程序不退出
	debug.FreeOSMemory()

	log.Println("done.")
	time.Sleep(1 * time.Hour) // 保持程序不退出
}
