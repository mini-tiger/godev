package main

import "fmt"
import (
	"math"
)

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 调用函数并返回最大值 */
	ret := max(a, b)

	fmt.Printf("max : %d \n", ret)
	lambda_ex()
	zhizhen()
}

/* 函数返回两个数的最大值 */
func max(num1 int, num2 int) int { //num1 num2 指定类型，  int 返回的类型
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
		return result
	} else {
		result = num2
		return result
	}
}

func lambda_ex() {

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 { //类似lambda 不用写函数名
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(8))

}

func zhizhen() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap(&a, &b)
	// var bb = a
	// a = b
	// b = bb

	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 使用*x方式 取x所保存的内存地址 上的值   */
	//temp 现在是  100
	fmt.Println(*x) // 100
	fmt.Println(x)  // 0xc042010248
	*x = *y         /* 将 y 值赋给 x,   ,*x打印出来是值，实际是对应内存地址 */
	*y = temp       /* 将 temp 值赋给 y ,*y取temp值的内存地址 */
}
