package main

import (
	"fmt"
)

//https://blog.csdn.net/qq_27068845/article/details/77407358
//可以通过  拆开切片 并行去重
func main() {
	m := []int{1, 2, 3, 1, 2, 2}
	fmt.Printf("切片长度:%d\n", len(m))
	mm := set1(&m) //mm := set(&m)  不同方法
	fmt.Printf("去重切片长度:%d\n", len(*mm))

}

func set(m *[]int) *[]int {
	tm := make(map[int]struct{}) //map init
	for _, k := range *m {
		tm[k] = struct{}{} //value 不使用
	}
	*m = []int{}
	for i, _ := range tm {
		fmt.Println(i)
		*m = append(*m, i) //重新赋值*m
	}
	return m

}

func set1(slc *[]int) *[]int {
	result := []int{}
	tempMap := map[int]struct{}{} // 存放不重复主键
	for _, e := range *slc {
		l := len(tempMap)
		tempMap[e] = struct{}{}
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return &result
}
