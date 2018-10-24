package main

import (
	"fmt"
	"container/list"
)

type M int
type N string

type I interface {
	S() string
}

func (self *M) S() string  {
	return fmt.Sprintf("int %d",*self)
}
func (self *N) S() string  {
	return fmt.Sprintf("string %s",*self)
}

func main()  {

	l:=list.New()
	m:=M(1)
	i:=I(&m)
	l.PushFront(i)
	//n:=N("1")
	//i=I(&n)

	//l.PushFront(i)

	l.PushFront([]int{1,2})
	l.PushFront(map[int]struct{}{1: struct{}{}})
	for e:=l.Front();e!=nil;e=e.Next(){
		ee:=e.Value
		fmt.Println(ee)

	}
	abc([]int{1,2}...)
}

func abc(a ...int)  {
	fmt.Println(a)
}