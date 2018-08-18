package main

import "fmt"

type age int
type name string

func (self *age) print_info()   {
	fmt.Printf("age: %v",*self)
}

func (self *name) print_info()   {
	fmt.Printf("name: %v",*self)
}
func main()  {
	a1:=age(1)
	fmt.Println(a1)
	a1.print_info()

	n1:=name("aa")
	fmt.Println(n1)
	n1.print_info()

}