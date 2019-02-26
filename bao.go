package main

import (
	"bytes"
	"io"
	"fmt"
	"os"
	"bufio"
	"time"
)

var ss=`
User Access Verification

Username: admin
Password: 
1F-bangong-1>en
Password: 
1F-bangong-1#show run | redirect tftp://10.240.81.86/10.240.80.11-20190226110$redirect tftp://10.240.81.86/10.240.80.11-201902261102         .cfg
!


















CORE-4506#show run | redirect tftp://10.240.81.86/10.240.80.254-201902261443.^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H$redirect tftp://10.240.81.86/10.240.80.254-201902261443.c         ^H^H^H^H^H^H^H^H^Hfg^M
!^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M
CORE-4506#Trying 10.240.80.11...^M
Connected to 10.240.80.11.^M
Escape character is '^]'.^M
User Access Verification^M
Username: admin^M
Password: ^M
1F-bangong-1>en^M
Password: ^M
1F-bangong-1#show run | redirect tftp://10.240.81.86/10.240.80.11-20190226144^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H$redirect tftp://10.240.81.86/10.240.80.11-201902261443         ^H^H^H^H^H^H^H^H^H.cfg^M
!^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M^M
1F-bangong-1#Trying 10.240.80.12...^M
Connected to 10.240.80.12.^M
Escape character is '^]'.^M
User Access Verification^M
Username: admin^M
Password: ^M
Connection closed by foreign host.   
`

func main() {
	//s,_ := StrTools.StrTrims(ss)
	//StrTrims(ss)
	DeleteBlankFile("c:\\run.log","c:\\11112.log")
}
func StrTrims(ss string) (d string,e error)  {
	dest := bytes.NewBuffer([]byte{})
	ssb := bytes.NewBuffer([]byte(ss))
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
		if 1 == len(str) || str == "\n" {
			fmt.Println(str)
			continue
		}
		//fmt.Print(str)
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
		fmt.Println(str=="\n")
		fmt.Println(str=="\r\n")
		fmt.Println(len(str))
		time.Sleep(2*time.Second)
		if err != nil {
			if err == io.EOF {
				fmt.Print("The file end is touched.")
				break
			} else {
				return err
			}
		}
		if 0 == len(str) || str == "\r\n" {
			fmt.Println(str)
			continue
		}
		//fmt.Print(str)
		destFile.WriteString(str)
	}
	return nil
}