package main

import (
	"os/exec"
	"strings"
	"fmt"
)

func getOSName() (cm string) {
	cmd := exec.Command("/bin/bash", "-c", "lsb_release -d")
	buf, _ := cmd.CombinedOutput()
	cmd.Run()
	cm = strings.Trim(string(buf), " ")
	cm = strings.Trim(cm,"Description:")
	cm = strings.Trim(cm, "\t")
	cm = strings.Trim(cm, "\r\n")
	return
}

func getOSVersion() (cm string) {
	cmd := exec.Command("/bin/bash", "-c", "lsb_release -r")
	buf, _ := cmd.CombinedOutput()
	cmd.Run()
	cm = strings.Trim(string(buf), " ")
	cm = strings.Trim(cm,"Release:")
	cm = strings.Trim(cm, "\t")
	cm = strings.Trim(cm, "\r\n")
	return
}
func main()  {
	fmt.Print(getOSName())
	fmt.Print(getOSVersion())
}