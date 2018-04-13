package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int)
	m1["age"] = 11 //添加，修改
	m1["long"] = 178
	fmt.Printf("%d \n", m1["age"])

	m2 := map[string]int{
		"age":  11,
		"long": 177,
	}

	fmt.Printf("%d \n", m2["age"])
	for s := range m2 {
		fmt.Println(s, m2[s])

	}

	del(&m1)
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
