package main

import (
	"fmt"
	"godev/g"
	flag "github.com/spf13/pflag"
)



func main()  {
	fmt.Println(g.Age)
	fmt.Println(g.Name)
	var ip= flag.IntP("flagname", "f", 1234, "help message")
	flag.Lookup("flagname").NoOptDefVal = "122" // nooptdefval 有标志项，没有添加数值
	flag.Parse()
	fmt.Println(*ip)
}