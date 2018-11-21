package main

import (
	"os"
	"fmt"
)

func IsFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {

		return false

	}
	return true
}
func main()  {
	if IsFile("/home/work/vsftpd/install.json "){
		fmt.Println(1)
	}
}