package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	d := "c:\\audio.log"
	fmt.Println(PathExists(d))
	fmt.Println(runtime.GOARCH)
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
