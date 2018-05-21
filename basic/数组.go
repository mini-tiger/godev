package main

import "fmt"

func main() {

	// b := []int{1, 2}   //type []int slice 长度空
	b := [2]int{1, 2}      // type [2]int， 长度固定的数组
	bl := [...]int{99: -1} //解释器判断长度，指定第99元素，其它不指定默认是0

	bb := [3]int{}
	fmt.Println(len(b), len(bl), bl[98], len(bb)) //2  100 0 3
	// fmt.Printf(b == balance)  // 不能比较，因为不同type

	/*
		[...]int(1,2) 效果不一样,没有指定长度，则根据给的数据, type [2]int
		[]int{}空这种方式，只能用在类型定义处,slice

	*/
	// fmt.Println(balance) // [0 0] 不指定默认元素是0
	fmt.Println("改变数组成员数值")
	b = [2]int{3, 4} // 类型不变，重新 赋值

	fmt.Printf("%T \n", b) //[2]int 是类型，
	fmt.Println(b[1])      //索引提取数据  2
	b[1] = 11              //索引改变数据
	fmt.Println(b[1])      //11

	fmt.Println("指向数组的指针")
	aa := new([2]int) //aa 返回指向数组的指针
	fmt.Println(aa)   //&[0 0]
	fmt.Println(*aa)  //[0 0]
	aa[0] = 1
	fmt.Println(*aa)

	fmt.Println("指针数组")
	mm, nn := 11, 22
	r := [2]*int{&mm, &nn}
	fmt.Println(r, *r[1])

	aaa := [...]int{1, 2}
	bbb := [...]int{0: 3, 1: 4}
	rr := [...]*[2]int{&aaa, &bbb}
	fmt.Println(rr, rr[1], *rr[1])

	two()    // 二维数组，数组参数传递
	maopao() //冒泡排序
}

func maopao() {
	ppp := [...]int{5, 2, 6, 3, 9}
	fmt.Println(ppp)
	cnt := len(ppp)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if ppp[i] > ppp[j] /*ppp[i] < ppp[j]*/ {
				temp := ppp[i] //只有在循环体中 ，可以重复 定义同一变量
				ppp[i] = ppp[j]
				ppp[j] = temp
			}
		}
	}
	fmt.Println(ppp)

}

func two() { //类似python [ [1，2]，[1，2]，[1，2]]
	/* 二维数组 -*/
	var a = [2][]int{{0, 0, 0}, {1, 2, 1, 1}}
	/*
	   [2]int 第一个元素{0, 0, 0}
	   [2]int 第二个元素{1，2，1,1}
	*/

	/* 输出数组元素 */
	num := len(a)
	for i := 0; i < num; i++ {
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
