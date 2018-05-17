package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Go编程"
	fmt.Println(len(s))

	fmt.Println(len([]byte{1, 2}))

	ss1 := []byte(s)
	ss2 := []rune(s)
	ss3 := utf8.RuneCount(ss1)
	ss4 := utf8.RuneCountInString(s)

	fmt.Println(ss1)
	fmt.Println(len(ss2))
	fmt.Println(ss3)
	fmt.Println(ss4)
}
