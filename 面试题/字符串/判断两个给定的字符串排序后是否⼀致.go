package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  判断两个给定的字符串排序后是否⼀致
 * @Version: 1.0.0
 * @Date: 2021/6/4 上午10:14
 */

func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))
	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isRegroup("ABC", "BCA"))
	fmt.Println(isRegroup("ABC", "BCAA"))
}
