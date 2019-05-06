package main

import (
	"fmt"
	"sort"
)


var m1 []int = make([]int,0)
func main () {

	m1=[]int{1,1,2,2,3,4,5,1}
	fmt.Println(m1)

	fmt.Println(sortgroupby(m1))
	}

func sortgroupby(source []int) [][]int {
	sort.Ints(m1)
	//fmt.Println(m1)
	m:=make([][]int,0)
	if len(source)==0{
		return m
	}
	i:=0
	j:=0

	for {
		if i>=len(source){
			break
		}
		for j = i + 1; j< len(source) && source[i] == source[j]; j++ {}

		m = append(m,source[i:j])
		i = j
	}
	return m
}

// todo 结构体按照某个元素 分组排序 https://blog.csdn.net/skh2015java/article/details/80197869
