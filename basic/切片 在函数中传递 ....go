package main

import "fmt"

func e() error {
	return fmt.Errorf("error: %s\n","this is error!!")
}

func main()  {
	fmt.Printf("%T,%v\n",e(),[]string{"a","b"})
	revier_slice([]string{"aa","bb"}...)
	//"后面加上三个点"
}

func revier_slice(s ...string)  {
	//接收一端可以不用  []string
	fmt.Printf("%T,%v",s,s)
}