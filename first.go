package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "Goodbye,")
	fmt.Println(s)
	s = strings.TrimPrefix(s, "Howdy,")
	fmt.Print("Hello" + s)
}
