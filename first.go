package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "7983* , SWC(C)"
	rex := regexp.MustCompile(`\(([a-zA-Z)]+)\)`)
	out := rex.FindAllStringSubmatch(str, -1)

	for _, i := range out {
		fmt.Println(i)
	}
}