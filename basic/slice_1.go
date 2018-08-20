package main

import "fmt"

var s1 [10]int
//一、有底层数组的 切片
//1. 左边被截取, cap = len      arr[2:]，中间截取 也属于左边截取
//2. 右边截取,CAP与底层数组一样   arr[:5]
//3. 超过CAP，都是每次增长自身CAP一倍
func main()  {
	// slice 初始化长度3，容量6
	s3:=make([]int ,3,6)
	print_info("s3",s3) //切片 len: 3,cap:6,value:[0 0 0]

	s1 = [10]int{0,1,2,3,4,5,6,7,8,9}
	s4 := s1[:] //s1是S3的底层数组，长度容量都和底层数组一样，10
	print_info("s4",s4)
	s4=append(s4,1) //添加一个到切片后，CAP一次性 增长一倍，len加上添加元素的个数
	print_info("s4",s4) //切片 len: 11,cap:20,value:[0 1 2 3 4 5 6 7 8 9 1]

	s22 := s1[2:] //左开右闭,从左边截取，cap = len
	print_info("s22",s22)
	s22=append(s22,1,2,3,4) //添加元素后，CAP超过底层数组，一次增长一倍自身CAP
	print_info("s22",s22)
	//
	s33 := s1[:8] // 从右边截取，CAP和底层数组一样
	print_info("s33",s33)
	s33 = append(s33,1,2,3,4,5) // //添加元素后，CAP超过底层数组，一次增长一倍自身CAP
	print_info("s33",s33)
	//原生切片
	fmt.Println("==================原生切片===================")
	s55 := []int{1,2,3}
	print_info("s55",s55)  //定义时len=cap
	s55 = append(s55,1,2,3,54,6,6,7)
	print_info("s55",s55)  //增长后 len一直与cap一样

	s66 := make([]int,2) //初始len cap为2。 {0，0}
	print_info("s66",s66)
	s66=append(s66,1,2,3,4,5,6,7)//增加 7个元素，每次增长cap 2，len 实际元素， 加完之后是 len 9  cap 10
	print_info("s66",s66)
	}
func print_info(name string,s ...interface{})  {
	switch s[0].(type){
		case []int:
			x:=s[0].([]int)
			fmt.Printf("切片%s, len: %d,cap:%d,value:%v\n",name,len(x),cap(x),x)

	}


}
