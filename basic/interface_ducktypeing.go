package main

import "fmt"

type student struct {
	name string
	age  int
}

type teacher struct {
	name string
	age  int
}

type human interface {
	name_str() string
}

func (self *student) name_str() string {
	return fmt.Sprintf("myname is %s,age is: %d\n", self.name, self.age)
}
func (self *teacher) name_str() string {
	return fmt.Sprintf("myname is %s,age is: %d\n", self.name, self.age)
}

func main() {
	var s1 student = student{name: "ssss", age: 11}
	var t1 teacher = teacher{name: "tttttt", age: 31}
	fmt.Println(s1.name_str())
	fmt.Println(t1.name_str())
	fmt.Println("==========================下面是接口==========================")
	arr_interface := []human{}

	arr_interface = append(arr_interface, &s1, &t1)
	for _, model_name := range arr_interface { // 类型都是 结构体的类型，接口本身只是起到统一的作用
		fmt.Printf("type:%T,value:%+v,name_str:%s\n", model_name, model_name, model_name.name_str())
	}

}
