package main

import (
	"bytes"
	. "fmt"
)

func main() {
	var a [10]byte = [10]byte{1, 2, 3}
	Printf("%d\n", a)
	// 将数组a转换为切片
	b := a[:]
	// 去除切片尾部的所有0
	c := bytes.TrimRight(b, "\x00")
	Printf("%d\n", c)

	sa := []byte("Golang")
	Println(string(sa[0:2])) //go
	Printf("%q \n", sa[0:2])

	Println(bytes.Count([]byte("cheese"), []byte("e"))) // 3
	Println(bytes.Count([]byte("five"), []byte("f")))   // 1

	Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	Println(bytes.HasSuffix(sa, []byte("ang")))

	Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a "))) //["" "man " "plan " "canal panama"]
	Printf("[%q]\n", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
}
