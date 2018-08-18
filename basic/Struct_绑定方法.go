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

func (self *Human) age_return() {
	fmt.Printf("name: %s, age:%d \n", self.name, self.age)
}

func age_return(self *Human) { //普通函数名 与  绑定的方法名可以重名
	fmt.Printf("name: %s, age:%d \n", self.name, self.age)
}
func main() {
	var h1 Human = Human{}
	h1.name = "张三"
	h1.age = 18
	h1.age_return()
	age_return(&h1)
}
