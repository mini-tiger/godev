package main

import "fmt"

type A struct {
	name string
}

type inter interface {
}

//空接口可以承载任意类型， map[string]interface{},  []interface{}
// struct 赋值给 空接口后，不能通过接口，调用 自身的属性与方法
// 使用接口断言 ，确定类型并 打印结构体属性

func variables(i interface{}) {

	switch i.(type) { //接口对象.(实际类型)

	case int:
		fmt.Println("type : int ,value:", i)
	case float64:
		fmt.Println("type : float64 ,value:", i)
	case A:
		fmt.Println("type : A ,value:", i)
	case []int:
		fmt.Println("type : []int ,value:", i)

	}

}
func main() {
	var a inter = 1
	var b inter = 1.1
	var c inter = A{"aa"}
	d := make([]int, 0, 10)
	d = append(d, 0)
	variables(a)
	variables(b)
	variables(c)
	variables(d)
}
