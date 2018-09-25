package main

import "fmt"

func main() {
	d := make([]int,5)
	d1 := make([]int,1)
	fmt.Println(d, cap(d))
	copy(d, []int{1, 2, 3})
	copy(d1, []int{1, 2, 3})
	fmt.Printf("copy后切片:%v\n",d) // copy后切片:[1 2 3 0 0]
	fmt.Printf("copy后切片:%v\n",d1) // copy后切片:[1] //todo copy不会超过目标切片CAP
	d=append(d,[]int{1,2,3}...)
	fmt.Printf("append后切片:%v\n",d) // append后切片:[1 2 3 0 0 1 2 3]

	a:=new(int)
	fmt.Printf("%T,%T\n",a,*a)
	fmt.Println(a)
	aa:=uintptr(*a)
	fmt.Println(aa)
	fmt.Printf("%T,%v\n",aa,aa)
	}
