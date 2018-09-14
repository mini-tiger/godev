package main

import (
	"fmt"
	"os"
)

func main() {
	//bytes.NewBuffer([]byte{99})
	s := 2 % 2
	fmt.Println(s)
	fmt.Println(string(os.PathSeparator))

}
