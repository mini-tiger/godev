package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("{Person name=%s, age=%d}", p.name, p.age)
}

func TestFormatPrint() {
	person := &Person{name: "andy", age: 12}
	fmt.Println(person)
}
func main()  {
	TestFormatPrint()
	//
	var s *name=&name{"ddd1"}
	fmt.Println(s)
}

type name struct {
	n string
}

func (self *name)String() string  {
	return fmt.Sprintf("this is n:%s\n",self.n)

}