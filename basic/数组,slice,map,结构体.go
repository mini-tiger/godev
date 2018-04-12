package main

import "fmt"

func main() {

	// b := []int{1, 2}   //type []int 长度空
	balance := [2]int{1, 2} // type [2]int， 长度固定
	bl := [...]int{99: -1}  //第99元素指定，其它不指定默认是0

	fmt.Println(len(balance), len(bl), bl[98]) //2  100 0
	// fmt.Printf(b == balance)  // 不能比较，因为不同type

	/*
		[...]int(1,2) 效果不一样,没有指定长度，则根据给的数据, type [2]int
		[]空这种方式，只能用在类型定义处

	*/
	// fmt.Println(balance) // [0 0] 不指定默认元素是0

	balance = [2]int{3, 4} // 类型不变，重新 赋值

	fmt.Printf("%T \n", balance) //[2]int 是类型，
	fmt.Println(balance[1])      //索引提取数据  2
	balance[1] = 11              //索引改变数据
	fmt.Println(balance[1])      //11

	type index int //索引通过常量代替
	const (
		a index = iota
		b
		c
		d
	)

	bb := []string{a: "aa", b: "bb"}
	fmt.Println(bb[a], bb[0]) // aa aa

	two() // 二维数组，数组参数传递
}

func two() { //类似python [ [1，2]，[1，2]，[1，2]]
	/* 二维数组 -*/
	var a = [2][]int{{0, 0, 0}, {1, 2, 1, 1}}
	/*
	   [2]int 第一个元素{0, 0, 0}
	   [2]int 第二个元素{1，2，1,1}
	*/

	/* 输出数组元素 */
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println(getAverage(a, len(a)))

	aa := [2]int{1, 2}
	fmt.Println(getAverage_pre(&aa))
}

func getAverage(arr [2][]int, size int) float32 {

	avg := float32(size) / float32(arr[1][1])

	return avg
}

func getAverage_pre(arr *[2]int) bool {
	fmt.Println(arr[0])
	return *arr == [2]int{1, 2} //true

}
