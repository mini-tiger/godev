package project_test //必须是文件夹名，也就是项目名

import (
	"fmt"
)

func Hello(s string) { //函数名必须大写
	fmt.Printf("Hello %v\n", s)
}

func Hi() {
	fmt.Printf("Hi test\n")

}
