package main

import "fmt"

type A struct {
	name string
	age  int
}

func (self *A) String() string {
	return fmt.Sprintf("myname is:%s, age is : %d\n", self.name, self.age)
}

func main() {
	var a1 A = A{"张三", 12}
	fmt.Println(&a1) //重写了   返回打印的方法， 由于上面定义 的方法是 指针，这里要传入地址

}
