package main

import "fmt"

type test struct {
	A []func()  int
}

var tt []test

func t1() int  {
	return 1
}
func t2() int {
	return 2
}
func t3() string {
	return "3"
}
func main()  {
	tt = []test{
		{A: []func() int{
			t1,
			t2,
		},
		},
	}
	fmt.Printf("%T\n",tt)
	for _,v:=range tt {
		for _,vv:=range v.A  {
			fmt.Printf("%T,%+v\n",vv,vv(  ))
		}

	}
}

