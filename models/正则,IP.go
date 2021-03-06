package main

import (
	"fmt"
	"regexp"
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
	fmt.Println("p1", i)
	i, err = regexp.MatchString(p, p2)
	fmt.Println(i)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	rr, _ := regexp.Compile(p)
	fmt.Println(rr.FindString(p1))

// xxx 找到括号内字符
	str := "7983* , SWC(C)"
	rex := regexp.MustCompile(`\(([a-zA-Z)]+)\)`)
	out := rex.FindAllStringSubmatch(str, -1)

	for _, i := range out {
		fmt.Println(i)
	}
}
