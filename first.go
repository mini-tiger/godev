package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(2)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(strings.TrimSpace(t2.Sub(t1).String()))
}
