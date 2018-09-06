package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main()  {
	//file:="c:/_tmp.txt"
	//temp,err:=os.Getwd()
	//
	//fmt.Println(strings.Split(temp,":"))
	////a:=filepath.VolumeName(file)
	// todo  windows 路径拼接 盘符要加 \\
	//file:=filepath.Join("c:\\","work","go-dev","log.txt")
	file:=filepath.Join("c:\\work","go-dev","log.txt")
	f,err:=os.Stat(file)
	fmt.Println(file)
	if os.IsNotExist(err){
		fmt.Println("no")
	}else{
		fmt.Println(f.Name()) //文件名
		fmt.Println(filepath.Dir(file)) //文件路径
		fmt.Println(filepath.VolumeName(file))//文件盘符
	}
	////aa:=fmt.Sprintf("%s%s","c:","1.txt")
	//fmt.Println(filepath.ToSlash(file))

	}