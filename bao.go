package main

import "fmt"

var ff []func()

func main() {

	ff = make([]func(), 0)
	ff = []func(){
		func() {
			fmt.Println("this is 1")

		},
		func() {
			fmt.Println("this is 2")
		},
		func() {
			fmt.Println("this is 3")
		},}

	for _,f:=range ff{
		f()
	}
}
