package main

import (
	"reflect"
	"fmt"
)

type A struct {
	AA  int
	AAA []int
}

func (a *A)Set(i int)  {
	a.AA=i
}


type B struct {
	BB  int
	BBB []int
	//A unsafe.Pointer
	A *A
}
func (b B)Set(i int)  {
	b.BB=i
}

func main() {
	var a A
	a.AA = 1
	a.AAA = []int{1, 2, 3}
	var b B
	b.BB = 2
	b.BBB = []int{4, 5, 6}
	gen(&a) // todo 通过反射函数的方法名，调用方法，传入实参，方法名必须大写，struct属性名必须大写
	fmt.Println(a)
	gen1(&b)  // todo 通过反射 操作结构体变量，必须传入指针
	fmt.Println(b)
	//aa:=(*A)(b.A)
	//fmt.Println(aa)

}

func gen(s interface{})  {
	sv:=reflect.ValueOf(s)

	method:=sv.MethodByName("Set") // 通过名称 找到方法
	vv:=[]reflect.Value{reflect.ValueOf(22)} // 转换转入的实参
	method.Call(vv)
}

func gen1(s interface{})  {
	sv:=reflect.ValueOf(s)
	// 方法一

	if sv.CanInterface(){
		// get
		sa:=sv.Elem()
		val := sa.FieldByName("BB").Int() // 提取
		fmt.Printf("N=%d\n", val) // prints
		// set
		mutable := reflect.ValueOf(s).Elem()
		mutable.FieldByName("BBB").Set(reflect.ValueOf([]int{1,1,1  })) // 赋值
		mutable.FieldByName("BB").SetInt(7)
		aa:=A{2,[]int{2,2,2}}
		//mutable.FieldByName("A").Set(reflect.ValueOf(unsafe.Pointer(&aa)))
		mutable.FieldByName("A").Set(reflect.ValueOf(&aa)) // todo 通过反射 属性名称，修改结构体属性，Set方法 支持所有类型
		//mutable.FieldByName("A").SetPointer(unsafe.Pointer(&aa)) // 修改 unsafe.Pointer 类型，Set方法的封装
	}



	//fmt.Printf("N=%d\n", ) // prints 7


	// 方法二
	//st:=sv.Type()
	//
	////st:=reflect.TypeOf(s)
	//
	////fmt.Println(st.Name(),st.String())
	//switch  {
	//case strings.Contains(st.String(),"A"):
	//	fmt.Println("A")
	//case strings.Contains(st.String(),"B"):
	//	if sv.Kind() == reflect.Ptr {
	//		tmp:=s.(*B)
	//		tmp.BB=33
	//	}
	//}

}