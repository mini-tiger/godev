package main

/*
string 不可变
btye可变

*/

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	s string = ";abjkh;lk;"
)
var ss = []string{"foo", "bar", "baz"}

func main() {

	fmt.Println(ImageFile("abc.jpg"))

}

func ImageFile(infile string) string {
	ext := filepath.Ext(infile)                                 // e.g., ".jpg", ".JPEG" 提取扩展名
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext // 删除扩展名 在 拼接字符串
	return outfile
}
