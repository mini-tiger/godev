package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  first
 * @Version: 1.0.0
 * @Date: 2021/9/10 下午2:56
 */

func main() {
	var arrs []string = []string{"1", "2", "3"}

	for _, v := range arrs {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(1 * time.Second)
}
