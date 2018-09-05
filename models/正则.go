package main

import (
	"regexp"
	"fmt"
)

func main() {
	ss := "abc中文"
	a, err := regexp.MatchString("^a.*中文",ss)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(a)
}
