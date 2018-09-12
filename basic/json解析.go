package main

import (
	"fmt"
	"os"
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
	os.Stdout.Write(b)
	fmt.Println()
	var cc interface{}
	json.Unmarshal(b, &cc) //解析JSON
	fmt.Printf("%T,%+v\n", cc, cc)

	//json 元素类型转换
	for _, v := range cc.([]interface{}) {
		switch v.(type) {
		case float64:
			fmt.Println("float64")
			v = v.(float64)
			fmt.Printf("%v\n", v)
		case string:
			fmt.Println("string")
			v = v.(string)
			fmt.Printf("%v\n", v)
		case []interface{}:
			//fmt.Println(v)

			strArray := make([]string, 0)
			for _, arg := range v.([]interface{}) {
				strArray = append(strArray, arg.(string))
			}
			fmt.Println(strArray)
		case map[string]interface{}:
			for k,vv:=range v.(map[string]interface{}){
				fmt.Println(k,vv)
			}

		default:
			fmt.Println("default")
			fmt.Printf("%T\n", v)
		}
		//fmt.Println(k,v)

	}
}
