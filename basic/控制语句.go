package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100*1.1 {
		/* if 条件语句为 true 执行 */
		if b == 200*2 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		} else {
			fmt.Printf("a 的值为 100 ， b 的值为 400\n")
		}
	} else {
		fmt.Printf("\n")
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
	switch_ex()
	switch_type_ex()
	select_ex()
}

func switch_ex() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
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

func select_ex() {
	var c1, c2, c3 chan int
	var i1, i2 int
	// fmt.Printf(c1, c2, c3)
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
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
}
