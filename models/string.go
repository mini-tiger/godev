package main

/*
string 不可变
btye可变

*/

import (
	. "fmt"
	"strings"
)

var (
	s string = ";abjkh;lk;"
)
var ss = []string{"foo", "bar", "baz"}

func main() {
	Println(strings.Join(ss, "\\")) //foo\bar\baz
	Println(strings.Count(s, "a"))  //统计包含
	Println(strings.Replace("oink oink oink", "k", "ky", 2))
	Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // -1 <0 ,不限次数替换
	Println(strings.HasPrefix("Gopher", "Go"))                    //判断开头结尾
	Println(strings.HasSuffix("Gopher", "Go"))
	Println(s, s[0:4])

	Println(strings.Index("chicken", "ken"))
	Println(strings.Index("chicken", "dmr"))  //查找索引 －1没有
	Println(strings.Index("go gopher", "go")) //rfind

	Printf("%q\n", strings.Split("a,b,c", ",")) //切分

	Println(strings.Trim(s, ";")) //删除前后
	Println(strings.Map(func(s rune) rune { return s + 1 }, "hello"))
	str := "hello roc"
	bytes := []byte(str)
	bytes[1] = 'a'
	str = string(bytes) //str == "hallo roc  改变字符串中的某些字符
}
