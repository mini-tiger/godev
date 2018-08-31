package funcs_ex

import (
	"os/exec"
	"bytes"
)



func Run_command(name string,args ...string) (string,error) {
	cmd := exec.Command(name,args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func Pipe_command(c string) (string,error) {
	cmd := exec.Command("sh", "-c", c)
	stdout_Stderr, err := cmd.CombinedOutput()
	if err != nil {
		return "",err
	}
	return string(stdout_Stderr),nil
}