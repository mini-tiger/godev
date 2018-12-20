package main

import (
	"github.com/pkg/errors"
	"fmt"
	"strings"
)

func main()  {
	err:=errors.New("123")
	fmt.Printf("%T,%s",err.Error(),err.Error())
	if strings.Contains(err.Error(),"123"){
		fmt.Println(1)
	}
}