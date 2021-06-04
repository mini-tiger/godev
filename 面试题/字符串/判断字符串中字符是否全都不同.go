package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: Tao Jun
 * @Description: 字符串
 * @File:  判断字符串中字符是否全都不同
 * @Version: 1.0.0
 * @Date: 2021/6/4 上午9:57
 */

//第⼀个是 ASCII字符 ， ASCII字符 字符⼀共有256个，其中128个是常
//⽤字符，可以在键盘上输⼊。128之后的是键盘上⽆法找到的。
//然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使⽤额外的储存结
//构，且字符串⼩于等于3000。

func main() {
	fmt.Println(isUniqueString2("Abcd123"))
	fmt.Println(isUniqueString2("AbcdA123"))
}

func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
