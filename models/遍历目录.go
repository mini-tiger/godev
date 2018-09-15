package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0]))) //当前文件父目录绝对路径
	fmt.Println(dir, err)
	d := filepath.Join("d:", string(os.PathSeparator), "work")
	log.Println(d)
	a, _ := filepath.Glob("d:\\work\\*dev\\*rc\\godev\\basic\\*.go")
	log.Println(a)
	var aa filepath.WalkFunc
	aa = func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(info)
		fmt.Println(err)
		return nil
	}
	filepath.Walk(d, aa)
}
