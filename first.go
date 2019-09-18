package main

import (
	"fmt"
	"strconv"
)

func main() {
	am:=new([]string)
	for i:=0;i<10;i++{
		*am=append(*am,strconv.Itoa(i))
	}
	fmt.Printf("%+v\n",(*am)[6:])
}