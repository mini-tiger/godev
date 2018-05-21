package main

import "fmt"
import (
	"time"
)

func main() {
	/* 定义局部变量 */
	// 	var a int = 100
	// 	var b int = 200

	// 	/* 判断条件 */
	// loop:
	// 	for i := 0; i < 10; i++ {
	// 		if i > 2 {
	// 			break loop
	// 		}
	// 		fmt.Printf("%v", i) //012
	// 	}

	// 	if a == 100*1.1 {
	// 		/* if 条件语句为 true 执行 */
	// 		if b == 200*2 {
	// 			 if 条件语句为 true 执行
	// 			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
	// 		} else { //else不能换行
	// 			fmt.Printf("a 的值为 100 ， b 的值为 400\n")
	// 		}
	// 	} else {
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("a 值为 : %d\n", a)
	// 	fmt.Printf("b 值为 : %d\n", b)

	if_ex()
	for_ex()
	switch_ex()
	switch_ex1()
	switch_type_ex()

	goto_break_continue()
	select_ex()
}

func if_ex() {
	if a := 1; a > 0 {
		fmt.Println("ifex >0 ")
	} else {
		fmt.Println("ifex <0 ")
	}

	var a string = "000"
	if a == "000" {
		fmt.Println("ifex 000")
	}

}

func for_ex() {
	for i := 0; i < 3; i++ {
		fmt.Printf("fox_ex:%d ", i)
	}
	fmt.Println()

	for a := 0; a <= 3; {
		a++
		fmt.Printf("fox_ex:%d ", a)
	}
	fmt.Println()
	a := 0
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Printf("fox_ex:%d ", a)
	}
	fmt.Println()

	aa := []int{1222, 21111}
	for _, a := range aa {
		tmp := &a //只有在循环体中 ，可以重复 定义同一变量
		fmt.Println(*tmp)

	}
}

func switch_ex1() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 91

	switch {
	case marks > 90:
		grade = "A"
	case marks > 80:
		grade = "B"
	default:
		grade = "D"
	}

	switch grade {
	case "A":
		fmt.Printf("优秀!\n")
	case "B", "C", "G": ////or 有一个true
		fmt.Printf("良好\n")
	case "D":
		fmt.Printf("及格\n")
	case "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
func switch_ex() {
	/* 定义局部变量 */
	var grade string = "D"
	// var marks int = 90

	switch marks := 90; {
	case marks >= 90:
		grade = "A"
		fallthrough //使用fallthrough强制执行后面的case代码，无论下一条结果是否为true
	case marks > 80:
		grade = "B" //每个都需要fallthough 往下
		fmt.Println(grade)
		// fallthrough
	case marks < 70:
		grade = "C"
	default:
		grade = "D"
	}

	// switch {
	// case grade == "A":
	// 	fmt.Printf("优秀!\n")
	// case grade == "B", grade == "C": //or 有一个true
	// 	fmt.Printf("良好\n")
	// case grade == "D":
	// 	fmt.Printf("及格\n")
	// case grade == "F":
	// 	fmt.Printf("不及格\n")
	// default:
	// 	fmt.Printf("差\n")
	// }
	// fmt.Printf("你的等级是 %s\n", grade)
}

//判断   x的类型
func switch_type_ex() {
	var x interface{}

	x = 1

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T \n", i)
	case int:
		fmt.Printf("x 是 int 型 \n")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

func goto_break_continue() {

	/*
		三个语法都可以配合标签使用
		标签名区分大小写，若不使用会造成编译错误
		Break与continue配合标签可用于多层循环的跳出
		Goto是调整执行位置，与其它2个语句配合标签的结果并不相同

	*/
	a := 0
label:
	for {
		a++
		if a > 3 {
			break label //break跳出后 不在执行 此处for
		}
	}
	fmt.Println(a) //4

label1:
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		// for {

		continue label1 //此label1效果与不加一样，每次跳过for的其中 一次循环
		// }
	}
	mm := 0
label2:
	for i := 0; i < 1; i++ {
		fmt.Println(i)
		mm++
		if mm > 3 {
			break label2
		}
		goto label2 //不加break 无限循环  重复执行label2下面的for
	}
}

func select_ex() {
	var c1 chan int = make(chan int)
	var c2 chan int = make(chan int)
	var c3 chan int = make(chan int)
	var i1, i2 int = 1, 1
	// fmt.Println(i1, i2)
	go func() { c1 <- 1 }()
	// time.Sleep(1 * time.Second)
	go func() {
		select {
		case <-c1:
			fmt.Println("received ", i1, " from c1")
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3): // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
		}
		// time.Sleep(3 * time.Second)
		// fmt.Printf("%d\n", <-c1)
	}()
	time.Sleep(3 * time.Second)
}
