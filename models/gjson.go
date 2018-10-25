package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"encoding/json"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {

	group :=[]ColorGroup{{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"}},
		{
			ID:     2,
			Name:   "green",
			Colors: []string{"Ab", "bc", "dd", "abbb"}},
	}
	b, err := json.Marshal(group) //打包 json
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
	m,ok := gjson.Parse(string(b)).Value().([]interface{}) //转换类型
	//fmt.Println(ok)
	if ok {
		fmt.Printf("%T\n", m)
		for k,v :=range m{
			fmt.Println(k,v)
		}

	}
	results:=gjson.Get(string(b),"last.name").Exists() //是否存在
	fmt.Println(results)
	r :=gjson.Get(string(b),"Name").Get("last") //级联提取,例子中没有
	fmt.Println(r)

	fmt.Println("===============unslice_json=============")
	unslice_json()


}
func unslice_json()  {
	group1 := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b1, err1 := json.Marshal(group1) //打包 json
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(string(b1))

	results:=gjson.GetMany(string(b1),"ID","Name") //一次性多个提取
	fmt.Println(results)


	value:=gjson.Get(string(b1),"Name")
	value1:=gjson.Get(string(b1),"Colors")
	fmt.Printf("%T,%v\n",value.String(),value.String())
	fmt.Printf("%T,%v\n",value1.Array(),value1.Array()) // todo 将提取的结果 转化为 数组，还可以MAP，
	for _,v := range value1.Array(){
		fmt.Printf("%T,%v\n",v.String(),v.String())
	}
}