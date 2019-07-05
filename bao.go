package main

import "fmt"


type A struct{
	AA []int
	BB map[int]int
	Ab AB
}


type AB struct{
	ABAB []int
}


func (a *A)getArry() []int {
	return a.AA
}

func (ab *AB)getArry() []int {
	return ab.ABAB
}


func main () {
	m1:=[]int{1,1,2,2,3,4,5,1}
	var a A=A{m1,map[int]int{1:1},AB{m1}}
	fmt.Println(a)
	fmt.Println(a.getArry())
	fmt.Println(a.Ab.getArry())
}
