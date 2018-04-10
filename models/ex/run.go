package ex

import (
	"fmt"
	"os"
	"path"
	// "path/filepath"
)

func Exist(file string) (string, bool) {
	if d, e := os.Getwd(); e == nil {
		f := path.Join(d, file)
		// fmt.Println(f)

		if _, err := os.Stat(f); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("文件: %s 不存在\n", f)
				return file, false
			}
		} else {
			fmt.Printf("文件: %s 存在\n", f)
			return file, true
		}
	}
	return file, false
}
