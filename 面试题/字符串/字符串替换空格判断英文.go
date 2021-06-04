package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  字符串替换空格判断英文
 * @Version: 1.0.0
 * @Date: 2021/6/4 上午10:32
 */

func check(s string) (string, bool) {
	ss := []rune(s)
	for _, v := range ss {
		if unicode.IsLetter(v) == false && string(v) != " " {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true // 将字符串中的空格全部替换为“%20”
}

func main() {
	fmt.Println(check("A B C a b "))
}
