package main

import "fmt"

var s []string=make([]string,0)

func main()  {
	fmt.Printf("%p\n",s)
	s=append(s,"1")
	fmt.Printf("%p\n",s)
	s=s[0:0]
	fmt.Printf("%p\n",s)
	s=append(s,"1")
	s=append(s,"2")
	fmt.Printf("%p\n",s)
	s=s[0:0]
	fmt.Printf("%p\n",s)
}