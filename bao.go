package main

import (
	"bytes"
	"io"
	"fmt"
	"os"
	"bufio"
)

var ss=`
str == "\r" 
CORE-4506#show run | redirect tftp://10.240.81.86/10.240.80.254-201902261443.$redirect tftp://10.240.81.86/10.240.80.254-201902261443.c         fg
!







122222222222   
`

func main() {
	//s,_ := StrTools.StrTrims(ss)
	//fmt.Println(StrTrims(ss))
	DeleteBlankFile("/root/run.log","1112.log")
}
func StrTrims(ss string) (d string,e error)  {
	dest := bytes.NewBuffer([]byte{})
	//ssb := bytes.NewBuffer([]byte(ss))
	src,_:=os.OpenFile("/root/run.log", os.O_RDONLY, 0666)
	ssb := bufio.NewReader(src)
	for {
		str, err := ssb.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//fmt.Print("The file end is touched.")
				break
			} else {
				//fmt.Println(err)
				e=err
				return
			}
		}

		if 0 == len(str) || str == "\r\n" {
			continue
		}
		if 1 == len(str) && str == "\n" {
			//fmt.Println(str)
			continue
		}
		if 1 == len(str) && str == "\r" {
			//fmt.Println(str)
			continue
		}
		dest.WriteString(str)
	}
	return dest.String(),nil
}


func DeleteBlankFile(srcFilePah string, destFilePath string) error {
	srcFile, err := os.OpenFile(srcFilePah, os.O_RDONLY, 0666)
	defer srcFile.Close()
	if err != nil {
		return err
	}
	srcReader := bufio.NewReader(srcFile)
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer destFile.Close()
	if err != nil {
		return err
	}

	for {
		str, _ := srcReader.ReadString('\r')
		fmt.Println(str)
		fmt.Println(str=="\r")
		//fmt.Println(str=="\n")
		//fmt.Println(str=="\r\n")
		fmt.Println(len(str))
		//time.Sleep(1*time.Second)
		if err != nil {
			if err == io.EOF {
				fmt.Print("The file end is touched.")
				break
			} else {
				return err
			}
		}
		if 0 == len(str) || str == "\r\n" {
			continue
		}
		if 1 == len(str) && (str == "\n" || str == "\r"){
			//fmt.Println(str)
			continue
		}
		//fmt.Print(str)
		destFile.WriteString(str)
	}
	return nil
}