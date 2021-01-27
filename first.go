package main

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"syscall"

	"os"
)

func main() {
	//err := ValidateContextDirectory("/root/1.sh")
	//if err != nil {
	//	fmt.Println(err)
	//}
	var filePath string = "/root/"
	var file os.FileInfo
	file, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(os.IsNotExist(err))
	//fmt.Println(os.IsPermission(err))

	//
	//lstat, _ := os.Lstat(filePath)
	//fmt.Println(lstat.IsDir(),lstat.Mode(),lstat.Sys())
	//
	//stat,_:=os.Stat(filePath)
	//fmt.Println(stat.Mode().Perm())
	//fmt.Println(stat.Mode().String())
	//fmt.Println(stat.Size())
	//fmt.Println(stat.ModTime())
	//fmt.Println(stat.Mode())
	fmt.Printf("%+v\n", file.Sys())
	rf := reflect.TypeOf(file.Sys())
	fmt.Println(rf.String())
	s := file.Sys().(*syscall.Stat_t)
	fmt.Printf("%+v\n", s)
	err = os.Chown(filePath, int(s.Uid), int(s.Gid))
	if os.IsPermission(err) {
		fmt.Println("not perm")
	}
}
func ValidateContextDirectory(srcPath string) error {
	var finalError error

	filepath.Walk(filepath.Join(srcPath, "."), func(filePath string, f os.FileInfo, err error) error {
		// skip this directory/file if it's not in the path, it won't get added to the context
		_, err = filepath.Rel(srcPath, filePath)
		if err != nil && os.IsPermission(err) {
			return nil
		}
		var err1 error
		if _, err1 := os.Stat(filePath); err1 != nil {
			finalError = fmt.Errorf("can't stat '%s'", filePath)
			//return err
		}
		if !os.IsPermission(err1) {
			fmt.Println(err1)
			return err1
		}

		// skip checking if symlinks point to non-existing files, such symlinks can be useful
		lstat, _ := os.Lstat(filePath)
		if lstat.Mode()&os.ModeSymlink == os.ModeSymlink {
			return err
		}

		if !f.IsDir() {
			currentFile, err := os.Open(filePath)
			if err != nil && os.IsPermission(err) {
				finalError = fmt.Errorf("no permission to read from '%s'", filePath)
				return err
			} else {
				currentFile.Close()
			}
		}
		return nil
	})
	return finalError
}
