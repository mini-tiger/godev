package main

import "fmt"

type age int //每个类型定义别名
type sex byte
type addr string
type name string

type Makelove struct {
	make int
}

type Human struct {
	age //不写属性名，默认类型就是属性名
	sex
	addr
	name
	*Makelove
	//聚合，继承其它结构体,一般直接用其它结构体的名字，也可以  abc *Makelove,使用其它名字
}

func (self *Makelove) make_str_ptr() (temp *int) {
	temp = &self.make
	return
}

func main() {
	m := &Makelove{7}
	fmt.Printf("makelove 原始value：%v, addr: %p \n", m.make_str_ptr, m.make_str_ptr()) //Makelove 结构体有make_str_ptr方法

	var h1 Human = Human{age: 1, sex: 1, addr: "this is", name: "alex", Makelove: m}
	//fmt.Printf("h1 type: %T , value： %v ,makelove value: %v ， makelove 绑定方法返回的指针地址: %p \n", h1, h1, h1.Makelove.make,h1.Makelove.make_str_ptr)
	// 子类 结构体 可以不用加 超类名，直接调用父类的属性  h1.Makelove.make 与 h1.make 相等
	// 不需要 调用 父类调用的字段叫  提升字段
	fmt.Printf("h1 type: %T , value： %+v ,makelove value: %v ， makelove 绑定方法返回的指针地址: %p \n", h1, h1, h1.make, h1.make_str_ptr)
	m.make = 10 //语法糖不用*
	fmt.Printf("h1 type: %T , value： %+v ,makelove value: %v ， makelove 绑定方法返回的指针地址: %p \n", h1, h1, h1.make, h1.make_str_ptr)
	fmt.Printf("makelove 更改后value：%v, addr: %p", m.make_str_ptr, m.make_str_ptr()) //Makelove 结构体有make_str_ptr方法

}
