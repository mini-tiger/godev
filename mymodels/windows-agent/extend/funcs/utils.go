package funcs_ex

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)

func Run_command(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func Pipe_command(c string) (string, error) {
	cmd := exec.Command("sh", "-c", c)
	stdout_Stderr, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(stdout_Stderr), nil
}

func Set(s []int) (temp []int) {
	t := map[int]struct{}{}
	for _, v := range s {
		t[v] = struct{}{}
	}

	for k, _ := range t {
		temp = append(temp, k)
	}
	return
}

//转换为中文
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
