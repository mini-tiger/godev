package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
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

	fmt.Println("=============use switch================")
	//json 元素类型转换
	for _, v := range cc.(map[string]interface{}) {
		switch v.(type) {
		case float64:
			fmt.Println("float64")
			v = v.(float64)
			fmt.Printf("%T,%v\n", v, v)
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
		default:
			fmt.Println("default")
			fmt.Printf("%T\n", v)
		}

	}
	fmt.Println("=============use reflect================")
	for _, v := range cc.(map[string]interface{}) {
		fmt.Println(Any(v))
	}

}

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	fmt.Println(v.Kind())
	switch v.Kind() {

	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)

	case reflect.Slice:
		strArray := make([]string, v.Len())
		//fmt.Printf("%v\n",v.Slice(0,v.Len()))
		type a interface{}
		var b a
		if v.CanInterface() { // 是否可以转换为接口
			b = v.Interface()
		} else {
			return "111111" // 否则返回错误
		}
		for i, vv := range b.([]interface{}) { //todo 再将 b 转换为[]interface
			strArray[i] = vv.(string) // 元素转换为字符串
		}

		return strings.Join(strArray, ",")
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 5, 32)

	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
