package main

import "fmt"

type A struct {
	name string
}

type inter interface {
}

//空接口可以承载任意类型， map[string]interface{},  []interface{}
// struct 赋值给 空接口后，不能通过接口，调用 自身的属性与方法
func main() {
	var a inter = 1
	var b inter = 1.1
	var c inter = A{"aa"}
	print_info(a)
	print_info(b)
	print_info(c)

	var mm map[string]interface{} = make(map[string]interface{})

	mm["aa"] = a
	mm["bb"] = b
	mm["cc"] = c
	print_info(mm)

	var slice1 []interface{} = make([]interface{}, 0, 10)
	slice1 = append(slice1, a)
	slice1 = append(slice1, b)
	slice1 = append(slice1, c)
	print_info(slice1)

	return_data(a)
	return_data(b)
	return_data(c)
}

func print_info(i inter) {
	fmt.Printf("type:%T,value:%+v\n", i, i)
}

func return_data(d ...interface{}) {
	fmt.Printf("type:%T,value:%+v\n", d, d)
}
