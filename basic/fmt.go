package main

import (
	. "fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	Printf("%v\n", p) // {1 2}
	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	Printf("%+v\n", p) // {x:1 y:2}
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	Printf("%#v\n", p) // main.point{x:1, y:2}

	var a [10]byte = [10]byte{1, 2, 3}
	Printf("%d %[1]T\n", a) // [1] 获取第一个参数
}
