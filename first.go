package main

import (
	"fmt"
	"regexp"
)

func main()  {
	s:="default//oracle/client"

	var dd =regexp.MustCompile("/.*")
	match :=dd.FindString(s)
	fmt.Println(match)

}