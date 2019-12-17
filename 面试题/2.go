package main

import "fmt"

type Test struct {
	Name string
}

func main() {
	list := make(map[string]Test)
	name := Test{"xiao"}
	list["name"] = name
	list["name"].Name = "Hello" // xxx map value不可寻址
	fmt.Println(list["name"])


	list1 := make(map[string]string)
	name1 := "xiao"
	list1["name"] = name1
	list1["name"] = "Hello"
	fmt.Println(list["name"])
}
