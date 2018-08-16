package main

import "fmt"

func main() {

	str := []int{1, 2, 3, 4, 5}
	fmt.Println(chose(&str, ji))
	fmt.Println(chose(&str, ou))

	str1 := []int{6, 7, 8, 9, 10}
	fmt.Println(chose(&str1, ji))
	fmt.Println(chose(&str1, ou))
	i, sum := 10, 0
	fmt.Println(*leijia(&i, &sum))
}

func leijia(i *int, sum *int) *int {
	if *i > 0 {
		*sum += *i
		*i--
		leijia(i, sum)
	} else {
		return sum
	}
	return sum
}

func ji(sl *[]int) []int {
	result := []int{}
	for _, x := range *sl {
		if x%2 != 0 {
			result = append(result, x)
		}
	}
	return result
}

func ou(sl *[]int) (result []int) { //默认返回result

	for _, x := range *sl {
		if x%2 == 0 {
			result = append(result, x)
		}
	}
	return
}

// myfunc 传入函数时 形参名，   func(*[]int) 要传入函数的参数， []int 函数返回的类型,  []int chose函数返回的类型
func chose(sl *[]int, myfunc func(*[]int) []int) []int {
	return myfunc(sl)
}

type myfunc func(*[]int) []int

func chose1(sl *[]int, myfun myfunc) []int {
	return myfun(sl)
}
