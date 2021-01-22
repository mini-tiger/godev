package main

import "fmt"

func GetNums(min int, max int, step int) (nums [][]int) {
	nums = make([][]int, 0)
	if min >= max {
		return
	}
	allnums := make([]int, 0)
	for i := min; i <= max; i++ {
		allnums = append(allnums, i)
	}
	for i := 0; i < len(allnums); i = i + step {
		var tmpnum []int
		if i+step >= len(allnums) {
			tmpnum = allnums[i:len(allnums)]
		} else {
			tmpnum = allnums[i : i+step]
		}
		nums = append(nums, tmpnum)
	}

	return nums
}
func main() {
	fmt.Println(GetNums(0, 10, 2))
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a[0:2])
	fmt.Println()
}
