package main

import (
	"fmt"
	"sort"
)

func main() {
	m22 := new(map[int]string)
	fmt.Println(*m22)
	// m22 = {1:"a"}
	fmt.Println(m22)
	m22 = &map[int]string{
		1: "a",
	}
	fmt.Println(*m22)
	b := map[int]string{1: "a"}
	fmt.Println(b)
	m1 := make(map[string]int)
	m1["age"] = 11 //添加，修改
	m1["long"] = 178
	fmt.Printf("%d \n", m1["age"])

	m2 := map[string]int{
		"age":  11,
		"long": 177,
	}

	fmt.Printf("%d \n", m2["age"])
	for s := range m2 { //循环
		fmt.Println(s, m2[s])

	}

	if v, ok := b[1]; ok { //判断
		fmt.Println(v)
	}

	del(&m1)
	key()   //key排序
	value() //value排序
}

func del(m1 *map[string]int) {
	fmt.Println((*m1)["age"]) //11

	delete(*m1, "age")        //删除
	fmt.Println((*m1)["age"]) //0 没有则是0

	m := make(map[string]int)
	fmt.Println(m["age"]) //0

	if age, ok := m["age"]; ok {
		fmt.Println(age, ok)

	} else {
		fmt.Println(age, ok) //0 false
	}

}

func key() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

func value() {
	m := make(map[string]int)
	m["a"] = 1
	m["d"] = 4
	m["c"] = 3
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++

	}
	fmt.Println(p.Len(), p.Less(0, 1))
	sort.Sort(p)
	fmt.Println(p)
}

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
