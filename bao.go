package main

import (
	"bytes"
	"io"
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
	"tjtools/StrTools"
)

var ss = `
Using CATALINA_BASE:   /root/tomcat7
Using CATALINA_HOME:   /root/tomcat7
Using CATALINA_TMPDIR: /root/tomcat7/temp
Using JRE_HOME:        /data/bkce/service/java
Using CLASSPATH:       /root/tomcat7/bin/bootstrap.jar:/root/tomcat7/bin/tomcat-juli.jar
stop tomcat success
一月 02, 2019 2:51:59 下午 org.apache.catalina.core.StandardService stopInternal
信息: Stopping service Catalina
一月 02, 2019 2:51:59 下午 org.apache.coyote.AbstractProtocol stop
信息: Stopping ProtocolHandler ["http-bio-8080"]
一月 02, 2019 2:51:59 下午 org.apache.coyote.AbstractProtocol stop
信息: Stopping ProtocolHandler ["ajp-bio-8009"]
一月 02, 2019 2:51:59 下午 org.apache.coyote.AbstractProtocol destroy
信息: Destroying ProtocolHandler ["http-bio-8080"]
一月 02, 2019 2:51:59 下午 org.apache.coyote.AbstractProtocol destroy
信息: Destroying ProtocolHandler ["ajp-bio-8009"] 
`

func main() {
	s,_ := StrTools.StrTrims(ss)

	fmt.Println(s)
	//DeleteBlankFile("/root/run.log","1112.log")
}
func StrTrims(ss string) (d string, e error) {
	dest := bytes.NewBuffer([]byte{})
	//ssb := bytes.NewBuffer([]byte(ss))
	src, _ := os.OpenFile("/root/run.log", os.O_RDONLY, 0666)
	ssb := bufio.NewReader(src)
	for {
		str, err := ssb.ReadString('\n')
		//fmt.Println()

		if err != nil {
			if err == io.EOF {
				//fmt.Print("The file end is touched.")
				break
			} else {
				//fmt.Println(err)
				e = err
				return
			}
		}

		str = strings.Replace(str, string(byte(13)), "", -1) // 去除所有win 回车
		str = strings.Replace(str, string(byte(8)), "", -1)  //去除所有退格

		//if 0 == len(str) || str == "\r\n" {
		//	continue
		//}
		//if 1 == len(str) && (str == "\n" || str == "\r") {
		//	continue
		//}

		re := regexp.MustCompile("\\s{2,}?")
		str = re.ReplaceAllString(str, "")

		sl := strings.Count(str, string(byte(10))) //超过2 个换行，保留一个，其它删除

		if sl < 1 {
			str = str + string(byte(10))
		}
		//fmt.Println(str)
		//fmt.Println(1)
		//fmt.Println([]byte(str))
		dest.WriteString(str)
	}
	return dest.String(), nil
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
		fmt.Println(str == "\r")
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
		if 1 == len(str) && (str == "\n" || str == "\r") {
			//fmt.Println(str)
			continue
		}
		//fmt.Print(str)
		destFile.WriteString(str)
	}
	return nil
}
