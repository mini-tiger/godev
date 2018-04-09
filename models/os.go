package main

import (
	"fmt"
	"os"
)

func main() {
	a, err := os.Open("1.txt")
	fmt.Println(*a, err)

	a, err = os.Create("1.txt") //recreate
	fmt.Println(*a, err)
}
