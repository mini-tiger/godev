package main

import (
	"errors"
)

import (
	"bufio"
	"fmt"
	"github.com/CodyGuo/win"
	"io"
	"os/exec"
	"strings"
	"time"
)

var (
	winExecError = map[uint32]string{
		0:  "The system is out of memory or resources.",
		2:  "The .exe file is invalid.",
		3:  "The specified file was not found.",
		11: "The specified path was not found.",
	}
)

func main() {
	//err := execRun("cmd /c start http://www.baidu.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	for {
		fmt.Println(ExecOs())
		time.Sleep(1 * time.Second)
	}

}

func ExecOs() string {
	temps := ""
	cmd_string := fmt.Sprintf("wmic BaseBoard get %s", "manufacturer")
	cmd := exec.Command("cmd", "/c", cmd_string)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.Contains(strings.ToLower(string(line)), "manufacturer") {
			continue
		}
		temps = strings.TrimSpace(string(line))
		temps = strings.Trim(temps, "\n")
		break
	}
	return temps
}

func execRun(cmd string) error {
	lpCmdLine := win.StringToBytePtr(cmd)
	// http://baike.baidu.com/link?url=51sQomXsIt6OlYEAV74YZ0JkHDd2GbmzXcKj_4H1R4ILXvQNf3MXIscKnAkSR93e7Fyns4iTmSatDycEbHrXzq
	ret := win.WinExec(lpCmdLine, win.SW_HIDE)
	if ret <= 31 {
		return errors.New(winExecError[ret])
	}

	return nil
}
