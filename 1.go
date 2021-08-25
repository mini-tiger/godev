package main

import "fmt"

func main() {
	a := "中文"
	s := []byte(a)
	fmt.Println(string(s[0:1]))
	rs := []rune(a)
	fmt.Println(string(rs[0:1]))

}
