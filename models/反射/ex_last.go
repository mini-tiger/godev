package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (self User) Hello(i string) {
	fmt.Println("Hello world.", self.Name, i)
}

func main() {
	u := User{1, "OK", 12}
	Info(u)

	m := Manager{User: User{1, "Maria", 12}, title: "123"}
	t := reflect.TypeOf(m)

	// 通过反射获得匿名字段
	fmt.Printf("%#v\n", t.Field(1)) //第一个索引位置的字段，0
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	//由于是嵌套，使用FieldByIndex([]int{0,0}) ,第一个0代表索引位置0的匿名字段,第2个0匿名字段中的第0个索引位置
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))

	//设置

	set(&m)        //修改
	fmt.Println(m) // {{1 tttttt 12} 123}
	u.Hello("tao") //直接调用 方法 Hello world. OK tao
	func_b(&u)     //调用 结构绑定的方法 Hello world. OK tao
}

func Info(o interface{}) { //空接口，匹配所有类型
	t := reflect.TypeOf(o) //TypeOf 获取类型
	fmt.Printf("%T,Type:%s\n", t, t.Name())

	if t.Kind() != reflect.Struct { //判断 t的类型
		panic("传入的结构体 类型错误")
	}

	v := reflect.ValueOf(o) // 结构字段信息
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { //所有字段
		// 结构的字段信息
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ { //所有绑定的方法
		// 结构的方法信息
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
func set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && v.Elem().CanSet() { //是否是指针,是否可以修改

		v = v.Elem()
		fmt.Println(v)
		v.FieldByName("Age").SetInt(99) //修改年龄

	} else {
		fmt.Println("XXX")
		return

	}
	f := v.FieldByName("Name") // 按照结构体中 元素名字获取对象
	if f.IsValid() {           //如果有Name这个字段
		fmt.Println("has Field : Name")
		f.SetString("tttttt")
	}

}

func func_b(o interface{}) {
	v := reflect.ValueOf(o)
	f := v.MethodByName("Hello")
	if f.IsValid() {
		args := []reflect.Value{reflect.ValueOf("tao")} //传入到方法的变量定义
		f.Call(args)                                    //Call方法调用 绑定的方法
	}
}
