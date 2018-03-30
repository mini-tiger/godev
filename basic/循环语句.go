package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	for_ex()
	go_ex()
}

// func main() {
// 	var C, c int //声明变量
// 	C = 1        /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
// A:
// 	for C < 100 {
// 		C++ //C=1不能写入for这里就不能写入
// 		for c = 2; c < C; c++ {
// 			if C%c == 0 {
// 				goto A //若发现因子则不是素数
// 			}
// 		}
// 		fmt.Println(C, "是素数")
// 	}
// }

func for_ex() {
	/* 定义局部变量 */

	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= i; j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

func go_ex() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
	// LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			continue
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
