package main

import (
	ft "fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func (user *User) Print() {
	ft.Println("reflect Print()")
}
func (user User) Print1() {
	ft.Println("reflect Print()")
}

func Reflect(inter interface{}) {
	t := reflect.TypeOf(inter) //从接口中获取结构的对象

	if k:=t.Kind();k!=reflect.Struct{//判断传入的是否是struce类型,而不是指针类型*User,指针类型报错
		ft.Println("type is not true")
		return
	}
	ft.Println("类型名称:", t.Name())
	v := reflect.ValueOf(inter) //从接口中获取结构的值

	for i := 0; i < t.NumField(); i++ { //遍历所包含的属性字段
		f := t.Field(i) //获取到字段
		val := v.Field(i).Interface()
		ft.Println("字段签名:", f.Type, " 字段名称:", f.Name, "  值:", val)
	}

	for i := 0; i < t.NumMethod(); i++ { //遍历所绑定的方法
		m := t.Method(i) //获取到方法
		ft.Println("方法名称:", m.Name, " 方法签名:", m.Type)
	}
}


func main()  {
	u:=User{1,"haha"}
	Reflect(u)
}