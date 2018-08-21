package main

import (
	"fmt"
	"os"
)

func main() {
	d, _ := os.Getwd()
	fmt.Println(d)
}
