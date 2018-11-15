package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	s, e := os.Open("D:\\putty")
	fmt.Println(e)
	f, e := s.Stat()
	fmt.Println(f.IsDir())
	fmt.Println(filepath.Dir("putty\\PAGEANT.EXE"))
}
