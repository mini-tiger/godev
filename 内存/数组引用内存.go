package main

import "fmt"

// todo 数组的Struct 是 第一个元素的内存地址 和 长度，组成的struct
func main() {
	s := make([]int, 3, 3)
	s1 := s
	s2 := s
	s[0] = 123
	s1[1] = 456
	s2[2] = 789 // s,s1,s2 指向同一块地址
	fmt.Println("原始切片：", s)
	fmt.Println("赋值切片：", s1)
	fmt.Println("赋值切片：", s2)
	fmt.Println("切片的长度、容量为：", len(s), cap(s))
	fmt.Printf("追加之前的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2) // xxx 地址一样

	fmt.Println("====执行append操作====")
	s2 = append(s2, 123)
	fmt.Printf("追加之后的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2) //xxx s2 超过了最大容量，会开辟新的内存地址，把数据复制过去
	s[0] = 888                                            // TODO：会修改s1、s2吗
	fmt.Println("原始切片s：", s)
	fmt.Println("赋值切片s1：", s1)
	fmt.Println("赋值切片s2：", s2)
	fmt.Println("以上 s2 已经不是最开始的s2地址，不会改变")
	s = append(s, 666) //xxx s 超过了最大容量，会开辟新的内存地址，把数据复制过去
	fmt.Println("原始切片s：", s)
	fmt.Println("赋值切片s1：", s1)
	fmt.Println("赋值切片s2：", s2)
	fmt.Printf("追加之后的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2) // 三个地址都 不一样
}
