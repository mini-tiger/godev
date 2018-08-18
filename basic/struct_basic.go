package main

import "fmt"

type Struct_1 struct {
	A string
	B int
}

func main() {

	//声明,必须实例化后才能使用
	var s1 Struct_1 = Struct_1{"s1", 2}

	s2 := Struct_1{ //悬挂
		A: "s2",
		B: 1,
	}

	var s3 *Struct_1 = &Struct_1{"s3", 3}

	s4 := &Struct_1{}
	(*s4).A = "s4"
	(*s4).B = 1

	var s5 Struct_1 //empty
	s6 := Struct_1{}

	s7 := new(Struct_1)

	fmt.Printf("s1: %v \n", s1)
	fmt.Printf("s2: %v \n", s2)
	fmt.Printf("s3：%v \n", s3)
	fmt.Printf("s4: %v \n", s4)
	fmt.Printf("s5：%v \n", s5)
	fmt.Printf("s6：%v \n", s6)
	fmt.Printf("s7：%v \n", s7)

	// 覆盖
	fmt.Println("====================覆盖==============================")
	s1.B = 3
	s1.A = "s1_1"
	fmt.Printf("s1: %v \n", s1)

	s2.A = "s2_1" // 重新设置其中一个
	fmt.Printf("s2: %v \n", s2)

	(*s3).A = "s3_1"
	s3.B = 11 //语法糖，不用写*
	fmt.Printf("s3: %v \n", s3)
}
