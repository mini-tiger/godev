package main

import "fmt"

type T interface {
	A() string
	B(int) int
}

type dd struct {
	A1 string
	B1 int
}

func (d *dd) A() string {
	return d.A1
}

func (d *dd) B(i int) int {
	d.B1 = i
	return d.B1
}

func main() {
	var d T = &dd{"1", 11}
	//var t T
	var d1 T = &dd{"2", 22} // 结构体 赋值接口
	d2:=T(&dd{"3", 33}) // 结构体 赋值接口

	s := []T{d, d1, d2}
	fmt.Printf("%+v\n",s)
	for i, v := range s {
		fmt.Println(v.A(), v.B(i))
	}
}
