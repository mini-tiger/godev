package main

import "fmt"

type Human struct {
	age  int
	name string
}

type student struct {
	class string
	human Human
}

func (self *Human) print_info() {
	fmt.Printf("myname %s,age:%d\n", self.name, self.age)
}

func (self *student) print_info() {
	fmt.Printf("myname %s,age:%d,class:%s\n", self.human.name, self.human.age, self.class)
}
func main() {
	var stu1 student = student{human: Human{age: 11, name: "san"}, class: "two"}
	fmt.Printf("%+v\n", stu1)
	stu1.print_info()       // 绑定 按就近原则，  绑定在 最外层子类上
	stu1.human.print_info() // 调用父类的方法

}
