package main

import (
	"fmt"
	"path/filepath"
)

func ClearMap() {
	tmpLock := make(chan struct{}, 0)
	_=cap(tmpLock)
}

func main() {
	f,e:=filepath.Glob("/home/go/GoDevEach/works/haifei/syncHtml/htmlData/*")
	fmt.Println(f,e)

}