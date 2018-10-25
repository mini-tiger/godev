package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"github.com/tidwall/gjson"
)

type s struct {
	A int               `json:"a"`
	B string            `json:"b"`
	S []string          `json:"s"`
	M map[string]string `json:"m"`
}

func main() {
	ss := s{1, "1", []string{"1", "2"}, map[string]string{"a": "b"}}
	jsondata, err := json.Marshal(ss)
	fmt.Println(string(jsondata))
	fmt.Println(err)
	var cc interface{}
	//ss1:=&s{}
	json.Unmarshal(jsondata, &cc)
	//jiexi(cc)

	//r:=gjson.Get(string(jsondata),"m")
	//fmt.Println(r.Index,r.Num,r.Raw,r.Type,r.Str)
	r:=gjson.ParseBytes(jsondata).Get("m")
	fmt.Println(r)
}
func jiexi1(v interface{}) {

		vk := reflect.ValueOf(v)
		switch vk.Kind() {
		case reflect.String:
			v1 := v.(string)
			fmt.Printf("string %s",v1)
		case reflect.Int:
			v1 := v.(int)
			fmt.Println(v1)
		case reflect.Slice:
			v1 := v.([]interface{})
			for _, v2 := range v1 {
				fmt.Println(v2)
			}
		case reflect.Map:
			v1 := v.(map[string]interface{})
			fmt.Println(v1)
		}

}
func jiexi(v interface{}) {
	for _, vv := range v.(map[string]interface{}) {
		vk := reflect.ValueOf(vv)
		switch vk.Kind() {
		case reflect.String:
			v1 := vv.(string)
			fmt.Println(v1)
		case reflect.Int:
			v1 := vv.(int)
			fmt.Println(v1)
		case reflect.Slice:
			v1 := vv.([]interface{})
			for _, v2 := range v1 {
				fmt.Println(v2)
			}
		case reflect.Map:
			v1 := vv.(map[string]interface{})
			fmt.Println(v1)
		}
	}
}
