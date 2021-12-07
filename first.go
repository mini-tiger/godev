package main

import (
	"fmt"
	"reflect"
	"strings"
)

type BaseData struct {
	// mysql
	DbUrl      string `yaml:"db_url" name:"数据库地址"`
	DbUser     string `yaml:"db_user" name:"数据库用户名"`
	DbPassWord string `yaml:"db_pass_word" name:"数据库密码"`
	DbName     string `yaml:"db_name" name:"数据库名"`
}

func main() {
	d := BaseData{
		DbUrl:      "url",
		DbUser:     "user",
		DbPassWord: "pw",
		DbName:     "name",
	}
	fmt.Println(Struct2Map(d, true))
}

func Struct2Map(d interface{}, lower bool) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for k := 0; k < t.NumField(); k++ {
		if lower {
			m[strings.ToLower(t.Field(k).Name)] = v.Field(k).Interface()
		} else {
			m[(t.Field(k).Name)] = v.Field(k).Interface()
		}
		//fmt.Println("name:", fmt.Sprintf("%+v", t.Field(k).Name),
		//	", value:", fmt.Sprintf("%v", v.Field(k).Interface()),
		//	", yaml:", t.Field(k).Tag.Get("yaml"))
	}
	return m
}
