package main

import "fmt"


func main() {

	//var b bool=bool(true)
	var c []int=make([]int,5,10)
	fmt.Println(len(c))
	var cc []int
	cc=append(cc,[]int{1,2,3,4,5,6,7,8}...)
	fmt.Println(len(cc),cap(cc))

	var ccc []int=make([]int,0)
	ccc=append(ccc,[]int{1,2,3,4,5,6,7,8}...)
	fmt.Println(len(ccc),cap(ccc))


}
