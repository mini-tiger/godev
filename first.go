package main

import (
	"fmt"
	"strconv"
)

const (
	ga = 10
	gb = "gb"
)

type nameer interface {
	call()
}

type pp struct {
	a int
}

func (p *pp) call() {

	fmt.Println(p.a)
}

func main() {
	var n nameer
	var p1 pp
	p1.a = 1
	n = &p1
	n.call()

	a, b := 1, 2
	fmt.Println(a, b)
	if check(a, b) {
		fmt.Printf("this is true \n")
	} else {
		fmt.Printf("this is false \n")
	}

	var arr = []int{1, 2, 3, 4, 5}
	arr2 := [][]int{{1, 2}, {1, 2}}
	xunhuan(arr, arr2)
	gg()

	c := "33"
	d := "ss"
	zhizhen(&c, &d)
	fmt.Println(c, d)

	//类型转换
	e := 3.0
	f := "4"

	i, err := strconv.Atoi(f)

	fmt.Println(int(e), i, err)

	fmt.Println(*new(int))
	fmt.Println("ab" < "abc")

}

func zhizhen(x *string, y *string) {
	fmt.Println(x, y, *x, *y)
	temp := *x
	*x = *y
	*y = temp

}

func check(a, b int) bool {
	if a > 1 {
		return true

	} else {
		return false
	}

}

func xunhuan(args []int, args2 [][]int) {
	//一维数组
	for i := 0; i < len(args); i++ {
		fmt.Printf("this is index %d ,value %d \n", i, args[i])

	}
	//二维数组
	for i := 0; i < len(args2); i++ {

		for j := 0; j < len(args2[i]); j++ {
			fmt.Print(args2[i][j])

		}
		fmt.Println()

	}

}

func gg() {
	if ga >= 10 {
		for i := 0; i < 10; i++ {
			fmt.Println(i)

		}
	} else {
		fmt.Println("小于10 ")
	}

}
