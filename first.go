package main

import (
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 0) //make 建立的是指针
	//mm:=&m
	m = map[string]interface{}{"a": "a", "b": "b"}

	fmt.Printf("第一次打印原始地址：%p,值：%v,保存原始地址的指针地址:%p\n", m, m, &m)

	Test1(&m)
	Test2(m)
	fmt.Printf("第二次打印原始地址：%p,值：%v,保存原始地址的指针地址:%p\n", m, m, &m)
	m["a"] = "aaa"
	fmt.Printf("第三次打印原始地址：%p,值：%v,保存原始地址的指针地址:%p\n", m, m, &m)
}

func Test1(m *map[string]interface{}) {
	fmt.Printf("Test1函数内第一次打印地址：%p,值%v\n", m, *m)
	(*m)["a"] = "aa"
	fmt.Printf("Test1函数内第二次打印地址%p,值%v\n", m, *m)
}

func Test2(m map[string]interface{}) {
	fmt.Printf("Test2函数内第一次打印地址%p,值%v,内存地址:%p\n", m, m, &m)
	m["a"] = "aaaa"
	fmt.Printf("Test2函数内第二次打印地址%p,值%v，内存地址:%p\n", m, m, &m)
}
