package main

import "fmt"

type MM struct {
	S string
}

func (self MM)set1(k string)  { // todo 如果不使用指针，在调用set1时， self代表MM这个类的一个COPY，不是当前MM类
	fmt.Println("通过SET1 形参 方法")
	self.S=k
}

func (self *MM)set2(k string)  {
	fmt.Println("通过SET2 指针 方法")
	self.S=k
}


func main()  {
	var m MM
	m.set1("aa") //todo  set1调用当前类的COPY
	fmt.Println(m)  // {}
	m.set2("bb") //todo  调用当前类
	fmt.Println(m) // {bb}

	fmt.Println("==============================")
	m1:=&MM{}	//todo  set1调用当前类的COPY
	m1.set1("cc") // {}
	fmt.Println(m1)
	m1.set2("dd")
	fmt.Println(m1)




}