package main

import "fmt"

type A struct {
	A1 string
	A2 int
}


func main()  {
	var s *struct{
		x [][31]byte
	}
	fmt.Println(s.x[1024])
	fmt.Println(reflect.)
}