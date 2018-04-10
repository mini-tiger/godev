package main

import (
	e "./ex"
	"fmt"
	"os"
	// "path"
	"path/filepath"
)

// func exist(file string) bool {
// 	if d, e := os.Getwd(); e == nil {
// 		f := path.Join(d, file)
// 		// fmt.Println(f)

// 		if _, err := os.Stat(f); err != nil {
// 			if os.IsNotExist(err) {
// 				fmt.Printf("文件: %s 不存在\n", f)
// 				return false
// 			}
// 		} else {
// 			fmt.Printf("文件: %s 存在\n", f)
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	// a,b=os.getwd()
	fmt.Println(os.Args[0])
	fmt.Println(os.Executable())
	var file string = "1.txt"
	// fmt.Println(file)
	if f, t := e.Exist(file); t {
		fmt.Printf("remove %s \n", f)
		os.Remove(f)
	} else {
		fmt.Printf("create %s \n", f)
		os.Create(f)
		fmt.Println(filepath.Abs(f))
	}
	// a, err := os.Open("1.txt")
	// fmt.Println(*a, err)

}
