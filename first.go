package main

import (
	"fmt"
	"strings"
)

func split(s *string, maxlen int) (*string) {
	if len(*s) > maxlen {
		(*s)=(*s)[0:maxlen-3]
		return s
	} else {
		return s
	}

}


func main() {
	a := "/ab/ca/bc11111ddddd"
	b:="a/ddd"
	c:="abc"
	//var index =
	aa:=a[strings.LastIndex(a, "/")+1:]
	fmt.Println(aa)

	//index =
	bb:=b[strings.LastIndex(b, "/")+1:]
	fmt.Println(bb)

	fmt.Println(strings.LastIndex(c, "/")+1)
	cc:=c[strings.LastIndex(c, "/")+1:]
	fmt.Println(cc)

}