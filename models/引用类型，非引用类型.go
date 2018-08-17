package main

import "fmt"

var a int =1
var arr [2]int = [2]int{1,2}

var sl1 []int = []int{1,2}

func main()  {
	int_type()
	arr_type()
	slice_type()

}

func slice_type(){
	fmt.Printf("切片 类型初始的值: %v \n",sl1)
	func(a *[]int) *[]int{
		(*a)[1] = 111
		return a
	}(&sl1)
	fmt.Printf("切片 类型 传递指针后的值: %v ，更改索引1位置的值 \n",sl1)

	func(a []int) []int{
		a[1] = 222
		return a
	}(sl1)
	fmt.Printf("切片 类型 不传递指针后的值: %v ，更改索引1位置的值，！！！变化 了，引用类型 \n",sl1)

}

func arr_type(){
	fmt.Printf("数组 类型初始的值: %v \n",arr)
	func(a *[2]int) *[2]int{
		(*a)[1] = 111
		return a
	}(&arr)
	fmt.Printf("数组 类型 传递指针后的值: %v ，更改索引1位置的值 \n",arr)

	func(a [2]int) [2]int{
		a[1] = 222
		return a
	}(arr)
	fmt.Printf("数组 类型 不传递指针后的值: %v ，更改索引1位置的值，无变化 \n",arr)

}


func int_type(){
	fmt.Printf("int 类型初始的值: %d \n",a)
	func(a *int) *int{
		*a=2
		return a
	}(&a)
	fmt.Printf("int 类型 传递指针后的值: %d ，加了1\n",a)

	func(a int) int{
		a=4
		return a
	}(a)
	fmt.Printf("int 类型 不传递指针后的值: %d  ,无变化 \n",a)
}
