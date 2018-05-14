package main

import (
	"fmt"
)

type namer interface {
	String() string
	s()
}

type nn struct {
	a int
}

func (n *nn) String() string {
	return "111111"

}

func (n *nn) s() {
	fmt.Println("sssss")

}

func main() {

	type Stringer interface { //接口有String()方法
		String() string
	}
	type sss interface { //接口有s()方法
		s()
	}

	var x namer
	x = &nn{a: 2} //x有两个方法 String() ,s()

	// fmt.Println(x)
	switch i := x.(type) { //两个方法都匹配，只显示第一个
	case Stringer: // 是否有Stringer包含的String()
		fmt.Println(i.String())
	case sss: // 是否有sss包含的s()
		i.s()
	// case int:
	// 	fmt.Printf("x 是 int 型 \n")
	// case float64:
	// 	fmt.Printf("x 是 float64 型")
	// case func(int) float64:
	// 	fmt.Printf("x 是 func(int) 型")
	// case bool, string:
	// 	fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
