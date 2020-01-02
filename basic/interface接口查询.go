package main

import (
	"fmt"
)

type R interface {
	Read()
}

type W interface {
	Write(name string)
}

type RW interface {
	R
	W
}

type log struct {
	name []string
	r    int
}

func (t *log) Read() {
	if len(t.name) > t.r {
		fmt.Println(t.name[t.r])
		t.r++
	} else {
		fmt.Println(len(t.name),"empty")
	}
}
func (t *log) Write(name string) {
	t.name = append(t.name, name)
	fmt.Println("wirte success.", t.name)

}
func main() {
	var r R = &log{}
	var w W = &log{}
	w.Write("write first")
	w.Write("write second")
	r.Read() // xxx  r与 w分别是两个不同的对象，所以打印没有值
	r.Read()
	val, ok := w.(RW) // xxx 转换为其它接口 ，调用方法
	if !ok {
		fmt.Println("hi")
	} else {
		val.Read()
		val.Read()
	}
}
