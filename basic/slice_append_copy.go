package main

import "fmt"

//var sl1 []int

func main() {
	sl1 := make([]int, 0)
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		sl1 = append(sl1, v)
		fmt.Printf("sl1 的地址%p ,cap: %d,value: %v\n", &sl1, cap(sl1), sl1)
	}

	return_sl(sl1...)
	fmt.Println("===============================append=================")

	var sl2 []int
	sl2 = append(sl2, 1, 2, 3)              //apend可以接收 单个多个加入
	sl2 = append(sl2, sl1...)               //其它slice加入
	sl2 = append(sl2, []int{1, 2, 3, 4}...) //匿名slice 打散加入
	fmt.Printf("sl2 : cap:%d,value:%v\n", cap(sl2), sl2)

	fmt.Println("=======================append==copy=删改===========")
	//sl3=make([]int,0)
	fmt.Printf("删除前地址：%p,value:%v\n", &sl2, sl2)
	sl2 = sl2[1:]
	fmt.Printf("删除后地址：%p,value:%v\n", &sl2, sl2)

	sl3 := make([]int, 0)
	copy(sl3, sl2[1:5])
	copy(sl3, sl2[6:]) //删除了每5个
	fmt.Printf("sl3地址：%p,sl2地址：%p\n", &sl3, &sl2)

}

func return_sl(s ...int) {
	fmt.Printf("%T,%v\n", s, s)
}
