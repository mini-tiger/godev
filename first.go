package main

import "fmt"

func split(s *string, maxlen int) (*string) {
	if len(*s) > maxlen {
		(*s)=(*s)[0:maxlen-3]
		return s
	} else {
		return s
	}

}


func main()  {
	a:="abcabc11111ddddd"
	fmt.Printf("%v,%v\n",&a,a)
	split(&a,10)
	fmt.Printf("%v,%v\n",&a,a)
}