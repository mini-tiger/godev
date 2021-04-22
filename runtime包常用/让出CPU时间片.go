package main

import (
	"fmt"
	"runtime"
)

/**
 * @Author: Tao Jun
 * @Description: runtime包常用
 * @File:  让出CPU时间片
 * @Version: 1.0.0
 * @Date: 2021/4/15 下午5:34
 */

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Go routine...", i)
		}
	}()

	for i := 0; i < 10; i++ {
		runtime.Gosched() // xxx 让出CPU时间片,其他goruntines 执行
		fmt.Println("main func...", i)
	}
}
