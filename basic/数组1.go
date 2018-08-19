package main

import "fmt"

var arr1 [3]int
var arr2 [3]string = [3]string{"a", "b"} //少定义 一个，默认空字符串
//var arr3 [...]int

func main() {
	arr1 = [3]int{1, 2} //少定义 一个，默认0
	fmt.Print(arr1, "\n")
	arr3 := [...]int{1, 2, 3}
	fmt.Printf("arr3 length: %d\n", len(arr3))
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3 索引是%d,value:%d\n", i, arr3[i])
	}

	for ii, v := range arr3 {
		fmt.Printf("arr3 索引是%d,value:%d \n", ii, v)
	}

	//二维数组
	arr4 := [2][3]int{{1, 2, 3}}
	fmt.Println(arr4)

}
