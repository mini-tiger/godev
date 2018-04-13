package main

import "fmt"

/*
slice 类似数组 部分数据的 别名

*/
func main() {

	// b := []int{1, 2}   //type []int 长度空
	// balance := [2]int{1, 2} // type [2]int， 长度固定
	bl := [...]int{1: 11, 5: -1} //元素指定，其它不指定默认是0
	// s := []int{}

	fmt.Printf("%v,%[1]T ,%d\n", bl, len(bl)) // [0 0 0 0 0 -1],[6]int,   0到5  6个元素

	bl1 := bl[1:3] //slice 指针指向第一个 引用数组的元素即 索引为1

	fmt.Printf("%v,%[1]T %d, %d \n", bl1, len(bl1), cap(bl1))
	// [11 0],[]int 2, 5，   长度2 (1,2)   容量5 (1,2,3,4,5) 容量包括底层数组 被引用第一个元素以及后面所有

	/*
		bl 是 bl1(slice)的底层数组，bl1 长度2,容量5

	*/

	fmt.Printf("%v\n", bl1[:4]) // [11 0 0 0]
	// 超过定义容量从底层数组中提取

	bl[1] = 111      //底层数组重新赋值
	fmt.Println(bl1) //slice 随之修改

	bl1 = bl[:]
	fmt.Printf("%v\n", bl1)

	create(&bl)
	appen()
	del()
}
func create(sl *[6]int) {
	fmt.Printf("%v,%[1]T \n", (*sl)[:]) //指针创建

	fmt.Printf("%20s \n", "------------")
	sl1 := make([]string, 4, 6) // 建立切片，长度4，容量6，6可以省略，则和长度相同
	fmt.Printf("%T,%d,%d \n", sl1, len(sl1), cap(sl1))

	sl2 := []int{1, 2} //[]必须空，长度容量2
	fmt.Printf("%T,%d,%d \n", sl2, len(sl2), cap(sl2))

	sl3 := sl2[:] // slice生成 slice
	sl2[1] = 3
	fmt.Println(sl2, sl3) //底层slice 改变，随之
	fmt.Printf("%T,%d,%d \n", sl3, len(sl3), cap(sl3))

	//使用底层数组也可以创建

}

func appen() {
	fmt.Printf("%20s \n", "------------")
	sl := []int{1, 2}
	sl = append(sl, 11)
	fmt.Println(sl)
	fmt.Printf("%T,%d,%d \n", sl, len(sl), cap(sl))

	sl = append(sl, 5, 6) //多个元素
	fmt.Println(sl)

	ss := make([]int, 10, 20)
	ss1 := []int{4, 5, 6}
	ssss := append(ss, ss1...) //追加切片 使用 ...
	fmt.Println(ssss)          // [1 2 3 4 5 6]
	ssss[0] = 10
	ssss[1] = 12
	fmt.Println(ss, ss1)

	data := make([]int, 10, 20)
	data[0] = 1
	data[1] = 2
	dataappend := make([]int, 10, 20)
	dataappend[0] = 1
	dataappend[1] = 2
	result := append(data, dataappend...) //data=append(..)
	result[0] = 99                        //len <=10 则 	result[0] = 99 会 影响源Slice
	result[11] = 98
	fmt.Println("length:", len(data), ":", data)
	fmt.Println("length:", len(result), ":", result)
	// fmt.Println("length:", len(dataappend), ":", dataappend)
}

func del() {
	slice := []int{1, 2, 3, 4}
	slice[0] = slice[len(slice)-1]    //将要删除的索引值 被最后一个 覆盖
	fmt.Println(slice[:len(slice)-1]) // 不取最后一个

	slice = []int{1, 2, 3, 4}
	index := 2
	copy(slice[index:], slice[index+1:]) //索引位置后一个 往前复制
	fmt.Println(slice[:len(slice)-1])

}
