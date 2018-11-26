package main

import "fmt"

func main()  {
	exp()
}
func exp()  {
	var ss []int=[]int{1,2,3}
	var s int =1
	send(ss,s)
}
func send(abc ...interface{})  {
	fmt.Printf("%T,%+v\n",abc,abc)
	fmt.Println("abc len:%d",len(abc))
	recv(abc...)
}
func recv(aaa  ...interface{})  {
	fmt.Printf("%T,%+v\n",aaa,aaa)
	fmt.Println("aaa len:%d",len(aaa))
}