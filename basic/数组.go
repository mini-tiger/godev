package main

import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	//通过索引改变值
	n[0] = 200
	fmt.Println(n[0])
	two()
	func_chuandi()
}

func two() { //类似python [ [1，2]，[1，2]，[1，2]]
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println(len(a))
}

func func_chuandi() {
	/* 数组长度为 5 */
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	fmt.Println(len(balance))
	avg = getAverage(balance, len(balance))

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
}
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
