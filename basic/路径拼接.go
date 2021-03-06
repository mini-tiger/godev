package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func file_info(file string) {
	f, err := os.Stat(file)
	fmt.Println(file)
	if os.IsNotExist(err) {
		fmt.Println("no")
	} else {
		fmt.Println(f.Name())                  //文件名
		fmt.Println(filepath.Dir(file))        //文件路径
		fmt.Println(filepath.VolumeName(file)) //文件盘符
	}
}

func main() {
	//file:="c:/_tmp.txt"
	//temp,err:=os.Getwd()
	//
	//fmt.Println(strings.Split(temp,":"))
	////a:=filepath.VolumeName(file)
	// todo  windows 路径拼接 开头盘符要加 或者 string(os.PathSeparator)
	os_type := runtime.GOOS
	if os_type == "windows" {
		//file:=filepath.Join("c:\\","work","go-dev","log.txt")
		//file := filepath.Join("c:\\work", "go-dev", "log.txt")
		file := filepath.Join("c:", string(os.PathSeparator), "go-dev", "log.txt")
		file_info(file)

	}
	////aa:=fmt.Sprintf("%s%s","c:","1.txt")
	//fmt.Println(filepath.ToSlash(file))

	if os_type == "linux" {
		//p:=os.Getenv("GOPATH")
		//fmt.Println(p)
		// todo linux 路径拼接 开头根目录要加 /
		//file := filepath.Join("/home/go", "src","godev", "数组引用内存.go")
		file := filepath.Join("/home", "go", "src", "godev", "数组引用内存.go")
		file_info(file)
	}

}
