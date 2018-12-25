package main

import (
	"fmt"
	"os"
	"runtime"
)

func getline() (file string, line int) {
	_, file, line, _ = runtime.Caller(1)
	return
}

func main() {
	//a, b, c, e := runtime.Caller(0)
	//fmt.Println(a, b, c, e)
	tt()
}

func tt() {
	_, err := os.Getwd()
	if err == nil {
		fmt.Println(getline())
	}
	fmt.Println(getline())
}
