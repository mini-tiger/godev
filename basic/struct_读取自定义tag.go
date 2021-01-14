package main

import (
	_ "encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age" valid:"1-100"`
}

type OtherStruct struct {
	Age int `valid:"20-300"`
}

func validateStruct(s interface{}) bool {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fieldTag := v.Type().Field(i).Tag.Get("valid")
		fieldName := v.Type().Field(i).Name
		fieldType := v.Field(i).Type()
		fieldValue := v.Field(i).Interface()

		if fieldTag == "" || fieldTag == "-" {
			continue
		}

		if fieldName == "Age" && fieldType.String() == "int" {
			val := fieldValue.(int)

			tmp := strings.Split(fieldTag, "-")
			var min, max int
			min, _ = strconv.Atoi(tmp[0])
			max, _ = strconv.Atoi(tmp[1])
			if val >= min && val <= max {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	person1 := Person{"tom", 12}
	if validateStruct(person1) {
		fmt.Printf("person 1: valid\n")
	} else {
		fmt.Printf("person 1: invalid\n")
	}

	person2 := Person{"jerry", 250}
	if validateStruct(person2) {
		fmt.Printf("person 2: valid\n")
	} else {
		fmt.Printf("person 2: invalid\n")
	}

	other1 := OtherStruct{12}
	if validateStruct(other1) {
		fmt.Printf("other 1: valid\n")
	} else {
		fmt.Printf("other 1: invalid\n")
	}

	other2 := OtherStruct{250}
	if validateStruct(other2) {
		fmt.Printf("other 2: valid\n")
	} else {
		fmt.Printf("other 2: invalid\n")
	}
}
