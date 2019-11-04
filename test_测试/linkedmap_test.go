package test

import (
	"fmt"
	"testing"
)

func Benchmark_Linkmap(b *testing.B) {
	a := NewLinkedMap()
	a.Put("a", 1)
	a.Put("b", 2)
	a.Put("c", 3)
	//fmt.Println(a.Max())
	//fmt.Printf("%+v\n", a.MData)
	//fmt.Printf("%+v\n", a.MLink)
	for _, key := range a.SortLinkMap() {
		//fmt.Println(key)
		if v, e := a.Get(key); e {
			fmt.Printf("key:%s,value:%v\n", key, v)
		}
	}

	for _, key := range a.SortLinkMap() {
		//fmt.Println(key)
		if v, e := a.Get(key); e {
			fmt.Printf("key:%s,value:%v\n", key, v)
		}
	}
}
