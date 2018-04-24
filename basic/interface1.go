package main

import (
	"bytes"
	"fmt"
	// "io"
	"os"
	// "path/filepath"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	// fmt.Println(111)
	return len(p), nil
}

//!-bytecounter

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!\n")
	b.WriteTo(os.Stdout)

	// fmt.Fprintf(new(bytes.Buffer), "%v\n", 1111)
	// new(bytes.Buffer).WriteTo(os.Stdout)
	fmt.Fprintf(os.Stdout, "%v\n", 1111)

	// Fprintf 第一个参数 io.Writer, 传入的参数要有 Write 方法(源码中)，os.Stdout（os.Stdout.Write）,bytes.Buffer 都有此方法

	/*
		func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	*/

	var c ByteCounter
	c = 1                    //c 初始化1
	c.Write([]byte("hello")) //6
	fmt.Println(c)

	c = 0                         // c 赋值0
	fmt.Fprintf(&c, "hello,name") //c 加到10,等价于	c.Write([]byte("hello,name"))
	//Fprintf 调用c.write(), "hello,name" 直接转换为byte类型
	fmt.Println(c)
}
