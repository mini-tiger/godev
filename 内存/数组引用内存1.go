package main

import "fmt"

// todo 数组的Struct 是 第一个元素的内存地址 和 长度，组成的struct
func main() {
	s := make([]int, 3, 6)
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
	s2 = append(s2, 123)                                  //eee 现在s2是 [888 456 789 123]
	fmt.Printf("追加之后的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2) //xxx 没有超过最大容量，地址一样
	s[0] = 888
	s = append(s, 666) //eee 现在s 数组第4个元素是666
	fmt.Println("原始切片s：", s)
	fmt.Println("赋值切片s1：", s1)                            // eee 现在s1 最大长度是3，所以看不到第4个元素
	fmt.Println("赋值切片s2：", s2)                            // 由于底层数组是s,所以s追加的数据覆盖了s2 eee 现在s2 [888 456 789 666]
	fmt.Printf("追加之后的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2) //

	// 以下 将s1的长度 扩展到4个，其它引用同一地址的也会变动
	s1 = append(s1, 111)
	fmt.Println("原始切片s：", s)
	fmt.Println("赋值切片s1：", s1)
	fmt.Println("赋值切片s2：", s2)
}
