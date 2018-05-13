package intset

import (
	"testing"
)

func Test_Add(t *testing.T) {
	var (
		x IntSet
	)
	x.Add(1)
	x.Add(144)
	x.Add(9)
	//fmt.Println(x.String())// 打印结果
	if x.String() != "{1 9 144}" {//
		t.Error("不相等")
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	//fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	//fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	//fmt.Println(x.String()) // "{1 9 42 144}"
	if x.String() !="{1 9 42 144}" {
		t.Error("UnionWith Fail")
	}

}

func TestIntSet_Has(t *testing.T) {
	var (
		x IntSet
	)
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if !x.Has(1){//
		t.Error(x.Has(1))
	}
	if x.Has(2){// FALSE 不显示
		t.Error("Has Fail")
	}
}

func BenchmarkIntSet_Add(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		var (
			x IntSet
		)
		x.Add(1)

	}
}