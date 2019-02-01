package main

import "fmt"

type diy int

func main() {
	var i diy
	i = 1
	fmt.Println(i)

	func(inter interface{}){
		switch inter.(type) {
		case int:
			fmt.Println("int")
		case diy:
			fmt.Println("diy")

		}
	}(i)

	var ii int
	ii=2
	i=diy(ii)
	fmt.Println(i)
}
