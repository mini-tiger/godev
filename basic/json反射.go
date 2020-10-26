package main

import (
	"fmt"
	"reflect"
)

// Product 商品信息
type Product1 struct {
	Name      string `json:"name,omitempty"`                  //tag
	ProductID int64  `json:"product_id,omitempty" bson:"pid"` //omitempy，可以在序列化的时候忽略0值或者空值
	// todo omitempty 当数值为 默认值时，不能 生成带有此字段的json
	Number   int `json:"-" test:"test"` //tag  标签，json不序列化 此项
	Price    float64
	IsOnSale bool `json:"is_on_sale,string"` // 序列化转换 字符串
	Bt       bool `json:"bt"`                //首字母必须大写

}

func main() {
	//var p *Product1=&Product1{}
	//s:=reflect.TypeOf(p).Elem()
	//for i:=0;i<s.NumField();i++{
	//	fmt.Println(s.Field(0).Tag)
	//}

	var p Product1
	tt(p)


}
func tt(p interface{}) {
	typeof := reflect.TypeOf(p)
	//fmt.Println(typeof)

	field := typeof.Field(0)
	fmt.Println(field)
	fmt.Println(field.Tag.Get("json"))

	field = typeof.Field(1)
	fmt.Println(field)
	fmt.Println(field.Tag.Get("bson"))
	field = typeof.Field(2)
	fmt.Println(field)
	fmt.Println(field.Tag.Get("test"))
}
