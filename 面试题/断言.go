package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  断言
 * @Version: 1.0.0
 * @Date: 2021/5/14 下午3:54
 */

var i int = 1
var ii interface{} = 1

func main() {

	switch ii.(type) {
	case int:
		fmt.Println("int")
	}

	switch i.(type) { // 只能是interface 才能断言
	case int:
		fmt.Println("int")
	}

}
