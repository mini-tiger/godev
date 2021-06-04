package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  翻转字符串
 * @Version: 1.0.0
 * @Date: 2021/6/4 上午10:02
 */

func main() {
	fmt.Println(reverString("ABCD"))
}
func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
		//fmt.Println(i,string(str))
	}
	return string(str), true

}
