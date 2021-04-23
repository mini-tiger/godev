package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := []string{"A", "B", "C", "D", "E", "F", "G"}
	str2 := make([]string, 6)
	for i := 0; i < 6; i++ {
		str2[i] = strconv.Itoa(i)
	}

	str2 = append(str2, str1...)
	fmt.Println(str1)
	fmt.Println(str2)

}
