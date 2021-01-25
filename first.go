package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "http://btbtdy3.com/down/34845-0-11.html"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`(.*)(\d{2,3}).html`)
	if reg1 == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1[0][2])
}
