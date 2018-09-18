package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Product 商品信息
type Jtest struct { //struct每项名字要首字母大写
	Name        string                 `json:"name,omitempty"` //tag
	Url         string                 `json:"url,omitempty"`  //omitempy，可以在序列化的时候忽略0值或者空值
	Pages       int                    `json:"page"`           //tag  标签，json不序列化 此项
	Price       float64                `json:"-"`              // //tag  标签，json不序列化 此项
	IsNonProfit bool                   `json:"isNonProfit"`    // 序列化转换 字符串
	Address     map[string]interface{} `json:"address"`
	Links       []interface{}          `json:"links"`
}

var ss string = `{
    "name": "BeJson",
    "url": "http://www.bejson.com",
    "page": 88,
    "isNonProfit": true,
    "address": {
        "street": "科技园路",
        "city": "江苏苏州",
        "country": "中国"
    },
    "links": [
        {
            "name": "Google",
            "url": "http://www.google.com"
        },
        {
            "name": "Baidu",
            "url": "http://www.baidu.com"
        },
        {
            "name": "SoSo",
            "url": "http://www.SoSo.com"
        }
    ]
}`

func struct_func() {
	fmt.Printf("\n\n\n")
	var pp Jtest
	// pp := &Jtest{}                       // 关联 struct
	ee := json.Unmarshal([]byte(ss), &pp) //解析

	fmt.Println(ee)         // nil
	fmt.Printf("%+v\n", pp) //{Xiao mi 6 0 10000 2499 true false}
	fmt.Println(pp.Address["city"])
	a := pp.Links[0]
	x := switch_type_ex(a)
	fmt.Println(x["name"])
}

type link1 struct {
	name string
	url  string
}

func main() {
	fanshe()
	struct_func()

}

func fanshe() {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(ss), &dat); err == nil {
		// fmt.Println(dat)
		tt := reflect.ValueOf(dat["isNonProfit"])
		fmt.Println(tt.Kind(), tt.Bool() == true)
		fmt.Println(dat["isNonProfit"] == true)

		// fmt.Println(dat["name"])
		// fmt.Println(dat["address"])
	} else {
		fmt.Println(err)
	}
	fmt.Println("解析links")
	ass1 := reflect.ValueOf(dat["links"])
	// ass2 := reflect.ValueOf(dat["address"])
	// fmt.Println(ass)
	fmt.Println(ass1, ass1.Kind(), ass1.Len())
	for i := 0; i < ass1.Len(); i++ {
		t := ass1.Index(i)
		// k := reflect.ValueOf(t)
		fmt.Printf("%s	类型是否为接口:%t\n", t, t.Kind() == reflect.Interface)
		inter := t.Interface()
		fmt.Printf("接口内容： %#v\n", inter)
		x := switch_type_ex(inter)
		fmt.Printf("name:%s url:%s\n", x["name"], x["url"])

		// 		p := &link1{} // 关联 struct

		// e := json.Unmarshal([]byte(*), p) //解析

	}
	fmt.Printf("\n\n\n")
	fmt.Println("解析address")
	ass2 := reflect.ValueOf(dat["address"])
	fmt.Println(ass2, ass2.Kind())
	for _, k := range ass2.MapKeys() {
		tmp := ass2.MapIndex(k)
		// vtmp := reflect.ValueOf(tmp)
		fmt.Printf("字段:%-16s 内容类型:%v,内容:%v\n", k.String(), tmp.Kind() == reflect.Interface, tmp /*tmp.Interface()*/)
	}

	//以下通过 反射方式直接 转换dat, 现在 只能解析一层json数据
	// ass2 := reflect.ValueOf(dat)
	// fmt.Println(ass2, ass2.Kind(), ass2.MapKeys())

	// for _, k := range ass2.MapKeys() {
	// 	tmp := ass2.MapIndex(k)
	// 	// vtmp := reflect.ValueOf(tmp)

	// 	fmt.Printf("字段:%-16s 内容类型:%v,内容:%v\n", k.String(), tmp.Kind() == reflect.Interface, tmp /*tmp.Interface()*/)

	// }
}

func switch_type_ex(x interface{}) map[string]interface{} {
	// var x interface{}

	// x = 1

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T \n", i)
	case int:
		fmt.Printf("x 是 int 型 \n")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	case map[string]interface{}:
		fmt.Println("匹配map[string]interface{}")
		return i
	default:
		fmt.Printf("未知型\n")
	}
	return make(map[string]interface{})
}
