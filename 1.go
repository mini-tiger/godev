package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  结构体
 * @Version: 1.0.0
 * @Date: 2021/5/17 上午11:26
 */

type Foo struct {
	bar string
}

func main() {
	//s1 := []Foo{
	//	{"A"},
	//	{"B"},
	//	{"C"},
	//}

	var s1 []Foo
	s1 = []Foo{{"A"}, {"B"}, {"C"}}

	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Printf("%+v\n", s1)
	fmt.Println(s2[0], s2[1], s2[2])
}
