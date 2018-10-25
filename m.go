package main

import (
	"os/exec"

	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"strings"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)

		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func main() {
	cmd_string := fmt.Sprintf("tracert -d -h 5 -w 15 -4 www.baidu.com")
	fmt.Println("cmd", "/c", cmd_string)

	cmd := exec.Command("cmd", "/c", cmd_string)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)

	}
	ss := ConvertByte2String(stdout, GB18030)
	//fmt.Println(ss)
	//fmt.Println("===================================")
	ss = strings.TrimSpace(ss)
	ss = strings.Trim(ss, "\n")
	//fmt.Println(strings.Index(ss,"路由:"))
	s_sub := "路由"
	if strings.Index(ss, "路由:") != -1 {
		s_sub = fmt.Sprintf("%s:", s_sub)
	}
	ss1 := strings.Split(ss, s_sub)

	ss2 := strings.Split(ss1[1], "跟踪完成。")
	//ss = strings.Trim(ss2[0], "\n")
	//ss = strings.Trim(ss2[0], "\r")
	ss = strings.TrimSpace(ss2[0])

	ss3 := strings.Split(ss, "\r\n")
	fmt.Println(len(ss3))
	for _, v := range ss3 {
		sl := strings.Fields(v)
		temp := sl[len(sl)-1]
		if strings.Contains(temp, "请求超时") {
			continue
		}
		fmt.Println(temp)
	}
}
