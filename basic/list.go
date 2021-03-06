package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l1 *list.List = list.New()
	l2 := list.New() //指针类型
	fmt.Println((l1).Len())
	a := l1.PushBack("1")
	fmt.Println(l1.Len())
	l1.PushBack("2")
	l1.Remove(a)
	fmt.Println(l1.Len())
	fmt.Println(l1.Back().Value, l1.Front().Value)

	printInfo1("l1", l1)
	printInfo1("l2", l2)

	l1.PushFront("123") //插入开始位置
	l1.PushFront(1234)
	//l1.PushBack("中文")

	l2.PushFront("456")
	han := l2.PushBack("中文")        //插入结尾位置
	l2.InsertBefore("Before2", han) ////插入han之前位置

	l2.Remove(han) //删除元素

	printInfo1("l1", l1)
	printInfo1("l2", l2)
	iteror(l1)
	iteror(l2)

}
func printInfo1(name string, l *list.List) {
	fmt.Printf("type: %T,value:%v,length:%d\n", l, l, l.Len())
}

func iteror(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {

		fmt.Printf("type: %T,value:%v\n", e, e.Value)
	}

}
