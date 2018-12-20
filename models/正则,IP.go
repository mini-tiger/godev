package main

import (
	"regexp"
	"fmt"
)

// http://c2pblog.sinaapp.com/archives/504
func main() {
	ss := "abc中文"
	a, err := regexp.MatchString("^a.*中文", ss)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(a)


	p := `(\d+)\.(\d+)\.(\d+)\.(\d+)`
	p1 := "192.168.1.1"
	p2 := "127.0.0.1"
	i, err := regexp.MatchString(p, p1)
	fmt.Println(i)
	i, err = regexp.MatchString(p, p2)
	fmt.Println(i)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	rr, _ := regexp.Compile(p)
	fmt.Println(rr.FindString(p1))
}
