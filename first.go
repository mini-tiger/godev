package main

import (
	"errors"
	"fmt"
	"gitee.com/taojun319/tjtools/file"
	"os"
	"path/filepath"
	"strings"
)

func GetFatherDir(sf string) (string, error) {
	if !file.IsExist(sf) {
		return "", errors.New(sf + "Not Exist")
	}
	dirList := strings.Split(sf, string(os.PathSeparator))
	if len(dirList) <= 1 {
		return sf, nil
	}
	return strings.Join(dirList[0:len(dirList)-1], string(os.PathSeparator)), nil
}

func main() {
	a := "D:\\work\\project-dev\\GoDevEach\\works\\haifei\\syncHtml\\HisData\\10.135.13.164_5day_Chinese.html"
	fmt.Println(filepath.Base(a))
	fmt.Println(filepath.Split(a))
	dir, filename := filepath.Split(a)
	fmt.Println(dir)
	fmt.Println(filepath.Dir(dir))
	fmt.Println(filepath.Dir(filename))
	fmt.Println(filepath.Base(dir))
	fmt.Println(strings.Split(a, string(os.PathSeparator)))
	fmt.Println(strings.Split(a, "D:\\work\\project-dev\\GoDevEach\\works\\haifei\\syncHtml\\HisData\\"))

	fmt.Println(GetFatherDir(a))
	if b, e := GetFatherDir(a); e == nil {
		fmt.Println(GetFatherDir(b))
	}

	fmt.Println(GetFatherDir("d:\\"))
	fmt.Println(filepath.Abs("d:\\"))

}
