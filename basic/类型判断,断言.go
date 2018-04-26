package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = new(bytes.Buffer)

	f := w.(*bytes.Buffer)
	fmt.Printf("%T,%t\n", f, f == new(bytes.Buffer))

	w = os.Stdout
	ff := w.(*os.File)
	fmt.Printf("%T,%t\n", ff, ff == os.Stdout)

	var ww io.Writer = os.Stdout
	a, ok := ww.(*os.File)
	fmt.Println(*a, ok)
	/*
	   *bytes.Buffer,false
	   *os.File,true
	   {0xc042056080} true
	*/
	s := "BrainWu"
	if v, ok := interface{}(s).(string); ok {
		fmt.Println(v) //BrainWu
	}

	// 	pd_interface()
	tt()
	/*
	第二种，如果断言的类型T是一个接口类型，类型断言x.(T)检查x的动态类型是否满足T接口。

	如果这个检查成功，则检查结果的接口值的动态类型和动态值不变，但是该接口值的类型被转换为接口类型T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。
	如果检查失败，接下来这个操作会抛出panic，除非用两个变量来接收检查结果，如：f, ok := w.(io.ReadWriter)

	*/
	type Element interface{}    //接口类型
	var e Element = 100         //第一种，如果断言的类型T是一个具体类型，类型断言x.(T)就检查x的动态类型是否和T的类型相同。
	fmt.Println(type_switch(e)) //100
}

func type_switch(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "null"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
	case string:
		return x //在调用 实际函数前面， any(x)
	default:
		panic(fmt.Sprintf("type: %T,value: %[1]v", x))

	}
	return "1"
}

/*												*/
type name interface {
	name11() string
}
type name1 interface {
	name22()
}

type n struct{}

func (nn n) name11() string {
	return "1"
}

func (nn n) name22() {
	fmt.Println("tt")
}

func tt() {
	//接口name,name1，方法都与 结构体n绑定
	var n1 name
	n1 = n{}

	if f, ok1 := n1.(n); ok1 {
		fmt.Printf("%T %s\n", f, f.name11())
	}
	if f, ok2 := n1.(name1); ok2 { //同一 结构体绑定的接口， 转换为另一 接口
		f.name22()

	}
}

/*
断言类型的语法：x.(T)，这里x表示一个接口的类型，T表示一个类型（也可为接口类型）。
一个类型断言检查一个接口对象x的动态类型是否和断言的类型T匹配。

类型断言分两种情况：
第一种，如果断言的类型T是一个具体类型，类型断言x.(T)就检查x的动态类型是否和T的类型相同。

如果这个检查成功了，类型断言的结果是一个类型为T的对象，该对象的值为接口变量x的动态值。换句话说，具体类型的类型断言从它的操作对象中获得具体的值。
如果检查失败，接下来这个操作会抛出panic，除非用两个变量来接收检查结果，如：f, ok := w.(*os.File)
第二种，如果断言的类型T是一个接口类型，类型断言x.(T)检查x的动态类型是否满足T接口。

如果这个检查成功，则检查结果的接口值的动态类型和动态值不变，但是该接口值的类型被转换为接口类型T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。
如果检查失败，接下来这个操作会抛出panic，除非用两个变量来接收检查结果，如：f, ok := w.(io.ReadWriter)

*/

type Tester interface {
	getName() string
}
type Tester2 interface {
	printName()
}

//===Person类型====
type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}
func (p Person) printName() {
	fmt.Println(p.name)
}

//============
func pd_interface() {
	var t Tester
	t = Person{"xiaohua"}
	check(t)
}
func check(t Tester) {
	//第一种情况
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("%T , %s\n", f, f.getName())
	}
	//第二种情况
	if t, ok2 := t.(Tester2); ok2 { //重用变量名t（无需重新声明）,都是interface类型
		check2(t) //若类型断言为true，则新的t被转型为Tester2接口类型，但其动态类型和动态值不变
	}
}
func check2(t Tester2) {
	t.printName()
}
