package main

import "fmt"

type Struct_1 struct {
	A   int
	age byte
}

func main() {

	st1 := Struct_1{1, 2}
	fmt.Printf("st1 type: %T, value: %v, address: %p \n", st1, st1, &st1)

	//深拷贝，就是复制一个数值
	st2 := st1
	fmt.Printf("st2 type: %T, value: %v, address: %p \n", st2, st2, &st2)

	fmt.Println("非引用类型 必须使用指针  改变其它引用者的数值")
	st2.A = 1111
	fmt.Printf("st2 type: %T, value: %v, address: %p \n", st2, st2, &st2)
	fmt.Printf("st1 type: %T, value: %v, address: %p \n", st1, st1, &st1)

	//浅拷贝， 复制地址，会随着改变
	//方法一
	st3 := &st1
	st3.A = 2222 //语法糖，不用写 *
	fmt.Printf("st3 type: %T, value: %v, address: %p \n", st3, st3, &st3)
	fmt.Printf("st1 type: %T, value: %v, address: %p \n", st1, st1, &st1)

	//方法二
	st4 := new(Struct_1) //创建指针类型
	st4 = &st1
	st4.A = 3333 //修改 任意，两个都有影响
	st1.age = 18
	fmt.Printf("st4 type: %T, value: %v, address: %p \n", st4, st4, &st4)
	fmt.Printf("st1 type: %T, value: %v, address: %p \n", st1, st1, &st1)
}
