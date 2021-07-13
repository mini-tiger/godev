package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  可变参数...
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午9:44
 */

func main() {
	args(1, 2)
	args(1, 2, 3)
	args([]int{1, 2}...)
	args()
}

func args(n ...int) {
	fmt.Println(n)
}
