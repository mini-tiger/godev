package main

import (
	"time"
	"fmt"
)

func main() {
	a:=time.Now()
	time.Sleep(time.Duration(3)*time.Second)
	fmt.Println(a.Second())
	fmt.Println(time.Now().Second()-a.Second())
}