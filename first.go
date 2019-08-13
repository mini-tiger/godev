package main

import "fmt"

func main()  {
	for i,v:=range []int{1,2,3}{
		fmt.Println(v,i)
	}
	a([]int{1,2,3}...)
}

func a(i ...int)  {
	fmt.Println(i)
}