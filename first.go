package main

import (
	"fmt"
	"sync"
)

type A struct {
	AA string
}

func main() {
	// 初始化一个pool
	pool := &sync.Pool{
		// 默认的返回值设置，不写这个参数，默认是nil
		New: func() interface{} {
			return &A{}
		},
	}

	// 看一下初始的值，这里是返回0，如果不设置New函数，默认返回nil
	init := pool.Get().(*A)
	init.AA="a"
	fmt.Println(init)

	// 设置一个参数1
	pool.Put(1)
	pool.Put(2)

	// 获取查看结果
	num := pool.Get()
	fmt.Println(num)

	// 再次获取，会发现，已经是空的了，只能返回默认的值。
	num = pool.Get()
	fmt.Println(num)
}
