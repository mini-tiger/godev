package main

import (
	"path/filepath"
	"os"
	"fmt"
	"runtime"
)

func file_info(file string)  {
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

func main()  {
	//file:="c:/_tmp.txt"
	//temp,err:=os.Getwd()
	//
	//fmt.Println(strings.Split(temp,":"))
	////a:=filepath.VolumeName(file)
	// todo  windows 路径拼接 盘符要加 \\
	os_type:=runtime.GOOS
	if os_type =="windows" {
		//file:=filepath.Join("c:\\","work","go-dev","log.txt")
		file := filepath.Join("c:\\work", "go-dev", "log.txt")
		file_info(file)

	}
	////aa:=fmt.Sprintf("%s%s","c:","1.txt")
	//fmt.Println(filepath.ToSlash(file))

	if os_type == "linux"{
		p:=os.Getenv("GOPATH")
		fmt.Println(p)
		file := filepath.Join(p, "src","godev", "bao.go")
		file_info(file)
	}

	}