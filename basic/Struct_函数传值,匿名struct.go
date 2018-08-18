package main

import (
	"fmt"
	"reflect"
)

type human struct {
	age  int
	name string
	sex  byte
}

func main() {
	var h1 human = human{
		age:  11,
		name: "张三",
		sex:  1,
	}
	fmt.Printf("h1 type:%T ,value: %v \n", h1, h1)
	change_sex(h1) //传递值 "传递值到函数内 不会改变"
	fmt.Printf("h1 type:%T ,value: %v \n", h1, h1)
	change_sex(&h1) //传递值 "传递指针地址到函数内 会改变"
	fmt.Printf("h1 type:%T ,value: %v \n", h1, h1)

	//匿名struct
	fmt.Println("======================匿名Struct==================")
	a := struct {
		size int
		old  int
	}{
		size: 1,
		old:  11,
	}
	fmt.Println(a.old, a.size)

}

func change_sex(h interface{}) {

	switch h.(type) {
	case int:
		fmt.Println("int")

	case human:
		var temp human = h.(human)
		temp.age = 12
		fmt.Println("human")

	case uintptr:
		fmt.Println("uintprt")
	case *human:
		var temp *human = h.(*human)
		temp.age = 12
		fmt.Println("prt_human")

	default:
		fmt.Println("no switch")

	}

	h_type := reflect.ValueOf(h).Kind()
	if fmt.Sprintf("%v", h_type) == "struct" {
		fmt.Printf("h 传入函数的type: %v \n", h_type)
	}
	if fmt.Sprintf("%v", h_type) == "ptr" {
		fmt.Printf("h 传入函数的type: %v \n", h_type)
	}

}
