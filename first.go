package main

import (
	"fmt"
	"gitee.com/taojun319/tjtools/nmap"
)

func main() {
	s := nmap.NewSafeMap()
	s.Put("a", 1)
	fmt.Println(s)
}
