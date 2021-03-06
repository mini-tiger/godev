package main

import "fmt"

// import (
// 	"math/rand"
// )

/*
slice 类似数组 部分数据的 别名
*/
func main() {

	// b := []int{1, 2}   //type []int 长度空
	// balance := [2]int{1, 2} // type [2]int， 长度固定
	bl := [...]int{1: 11, 5: -1} //元素指定，其它不指定默认是0
	// s := []int{}

	fmt.Printf("%v,%[1]T ,%d\n", bl, len(bl)) // [0 0 0 0 0 -1],[6]int,   0到5  6个元素

	// bl1 := bl[1:3]                                         //slice 指针指向数组bl， 引用数组的元素即 索引为1至2，左开右闭，
	fmt.Printf("bl[1:]长度:%d\n", len(bl[1:]))               //索引 1,2,3,4,5
	fmt.Printf("bl[1:len(bl)]长度:%d\n", len(bl[1:len(bl)])) //索引 1,2,3,4,5
	fmt.Printf("bl[:4]长度:%d\n", len(bl[:4]))               // 索引 0，1，2，3

	slice1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bl2 := slice1[2:5]
	fmt.Printf("bl2: %v,len(): %d,cap(): %d \n", bl2, len(bl2), cap(bl2))
	// len()=3, 索引：234
	//cap()=8 索引：23456789 容量包括底层数组 被引用第一个元素以及后面所有

	// 超过定义容量从底层数组中提取
	fmt.Printf("超过slice范围的,从底层数组提取 bl2[:8]:%d\n", bl2[:8])
	//bl2 引用的第一个元素是数组索引2的位置，从2往后8个元素，[2 3 4 5 6 7 8 9]

	slice1[2], slice1[3] = 1, 1             //底层数组重新赋值
	fmt.Println("底层数组重新赋值,slice 随之修改", bl2) //slice 随之修改

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	a1 := a[1:3]
	a2 := a[2:5]
	fmt.Println(string(a1), string(a2))
	a[2] = 'Z'                          // 切片再切片, 修改底层切片,随之修改
	a2[1] = 'X'                         // 修改 生成出来的切片，所对应其它切片 也会修改
	fmt.Println(string(a1), string(a2)) //bZ ZXe

	a = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	a3 := a[1:3]
	a4 := a[2:5]
	// fmt.Println(string(a3), string(a4))
	a3 = append(a3, '1', '2', '3', '4', '5', '6', '7', '8', '9') //生成的切片，突破容量，重新分配内存
	a[2] = 'Z'
	fmt.Println(string(a3), string(a4)) //bc123456789 Zde 只有内存有映射关系的才会被修改

	s1 := []int{1, 2, 3}
	for index, v := range s1 { //循环
		fmt.Println(index, v)
	}

	create(&bl)
	appen()
	del_copy()

}
func create(sl *[6]int) {
	fmt.Printf("%20s \n", "------------")
	fmt.Printf("%v,%[1]T \n", (*sl)[:]) //指针创建

	a := 1
	sp1 := []*int{&a, &a} //指针数组
	fmt.Println(sp1)

	sl1 := make([]string, 1, 2) // 建立切片，长度4，容量6，6可以省略，则和长度相同
	fmt.Printf("%T,%d,%d \n", sl1, len(sl1), cap(sl1))

	//每当append到了容量最大值，自动扩展内存容量之前的两倍
	for i := 0; i < 6; i++ {
		sl1 = append(sl1, string(i))
		fmt.Printf("内存：%p第%d次追加,长度:%d,容量:%d \n", sl1, i+1, len(sl1), cap(sl1))

	}

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

func del_copy() {
	slice := []int{1, 2, 3, 4}
	slice[0] = slice[len(slice)-1]    //将要删除的索引值 被最后一个 覆盖
	fmt.Println(slice[:len(slice)-1]) // 不取最后一个

	slice = []int{1, 2, 3, 4}
	index := 2
	copy(slice[index:], slice[index+1:]) //索引位置后一个 往前复制
	fmt.Println(slice[:len(slice)-1])

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{0, 1}
	copy(s1, s2) //[0 1 3 4 5], s2 索引位置 覆盖到s1
	fmt.Println(s1)

	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{0, 1}
	copy(s1[4:5], s2[0:1]) //[1 2 3 4 0], s2 索引位置 覆盖到s1
	fmt.Println(s1)

	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	copy(s1, s2) //[0 1 2 3 4] 覆盖不会超过 s1最大长度
	fmt.Println(s1)
}
