package set

import (
	"testing"
)

func BenchmarkSet1(b *testing.B) {
	for i:=0;i<b.N ;i++ {

		m := []int{1, 2, 3, 1, 2, 2}
		//fmt.Printf("切片长度:%d\n", len(m))
		Set(&m) //mm := set(&m)  不同方法
		//fmt.Printf("去重切片长度:%d\n", len(*mm))
	}
}

func bench(b *testing.B,size int)  {
	b.N=size
	for i:=0;i<b.N ;i++ {
		m := []int{1, 2, 3, 1, 2, 2}
		//fmt.Printf("切片长度:%d\n", len(m))
		Set(&m) //mm := set(&m)  不同方法
		//fmt.Printf("去重切片长度:%d\n", len(*mm))
	}

}

func BenchmarkSet100000(b *testing.B) {
	bench(b,100000)
}