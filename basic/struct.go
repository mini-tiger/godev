package main

import (
	"fmt"
)

type Point struct { //定义,名字，变量名要大写，否则不能导出
	X, Y int //tag 用在json
	S    string
}

func main() {

	pp1 := Point{
		X: 11, //逗号 注意
		S: "1",
	}
	fmt.Println(pp1) //{11 0 1}

	pp2 := &struct { //匿名
		a int
		b string
	}{
		a: 1,
		b: "string",
	}
	fmt.Println(*pp2)

	pp3 := &struct { //匿名
		a int
		b string
	}{
		1, //按照顺序，不用写元素名
		"string",
	}
	fmt.Println(*pp3)

	p2 := &Point{} //初始化空值
	p2.S = "6"     //赋值1
	p2.X = 4
	p2.Y = 5
	fmt.Println(*p2) //{4 5 6}

	var p1 Point
	p1.S = "6" //赋值2
	p1.X = 4
	p1.Y = 5
	fmt.Println(p1) // {4 5 6}

	p := Point{1, 2, "3"}    //赋值3
	fmt.Println(p, p.X, p.S) //{1 2 3} 1 3

	var pp Point
	pp = Point{4, 5, "6"}
	fmt.Printf("%v %v %[1]T %[2]T\n", pp, pp.S) //{4 5 6} 6 main.Point string

	var ppp *Point = &pp //指针
	fmt.Println(*ppp)    //{4,5,6}

	pp.S = "change"       //改变， 影响内存引用
	fmt.Println(pp, *ppp) //{4 5 change} {4 5 change}

	duibi(&pp, &p1) //对比
	qintao()        //嵌套
}
func duibi(px *Point, px1 *Point) {
	fmt.Println((*px), (*px1))               //change
	fmt.Printf("%t \n", (*px).X == (*px1).X) //true
	fmt.Printf("%t \n", *px == *px1)         //false
	(*px1).S = "change"
	fmt.Printf("%t \n", *px == *px1) //true
}

func qintao() {
	type Point1 struct {
		Point //嵌套全局 Point
		M     int
	}

	pp := Point1{Point: Point{1, 2, "3"}, M: 3} //嵌套 初始化与p 效果一样

	p := Point1{Point{1, 2, "3"}, 3} //初始化
	p.M = 1
	fmt.Println(pp, p) //{{1 2 3} 3} {{1 2 3} 1}

	var p1 Point1  //初始化
	p1.S = "6"     //赋值
	p1.Point.X = 4 //等价于 p1.X=4
	p1.Y = 5
	p1.M = 7
	fmt.Println(p1) //{{4 5 6} 7}

}
