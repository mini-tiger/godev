package main

import (
	"fmt"
	"sync/atomic"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  mapBasic
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午9:29
 */

func main() {

	aa := make(map[string]int)
	delete(aa, "h")      // 什么都没有发生
	fmt.Println(aa["h"]) // 不存在， int默认值0

	ms()

	mloop()
}

type Math struct {
	x, y int
}

var m = map[string]Math{ // xxx 这里返回指针 则MAP.VALUE 可以寻址
	"foo": Math{2, 3},
}

func ms() {
	//m["foo"].x = 4  // 对于类似 X = Y 的 赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X ，但 go 中的 map 的 value 本身是不可寻址 的。

	fmt.Println(m["foo"].x)
}

func mloop() {
	fmt.Println("================mloop================")
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	var count int64 = 0
	for k, v := range m { //map是指针，副本也是指向源MAP的指针
		if count == 0 {
			delete(m, "A")
		}
		atomic.AddInt64(&count, 1)
		fmt.Println(k, v)
	}
	fmt.Println(count) //xxx 这里是 2 或者3,  MAP循环无顺序，第一次循环到A 则是3
}
