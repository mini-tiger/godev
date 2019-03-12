package main

import "fmt"

func main() {
	bbb()

}
func bbb() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	abc()
}
func abc() {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)

	}()
	panic(111)
}
