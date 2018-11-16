package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"fmt"
	"os/exec"
	"bytes"
)


type Monitor struct {
	Cmd string `json:"cmd"`
	Return int `json:"return"`
}

type InstallAll struct {
	Interval    int          `json:"interval"`
	SkipErr     bool         `json:"skipErr"`
	InstallStep map[int]interface{} `json:"install_step"`
	StartStep   map[int]interface{} `json:"start_step"`

}

type Install struct {
	InstallAll InstallAll `json:"install_all"`
	Monitor Monitor `json:"monitor"`
}


func UnJson(jb []byte) (err error, dl *Install) {
	err = json.Unmarshal(jb, &dl) //解析
	if err != nil {
		log.Println("解析json失败")
		return err, dl
	}
	return nil, dl
}

func main() {
	f, _ := ioutil.ReadFile("C:\\work\\go-dev\\src\\godev\\mymodels\\运行系统命令\\install.json")
	err, dl := UnJson(f) // 解析Json

	if err != nil {
		fmt.Println(1, err)
	}
	fmt.Println(dl)
	fmt.Println(dl.Monitor)
	fmt.Println(dl.InstallAll.InstallStep)


	//cinstall := make(chan struct{},0)
	//cstart := make(chan struct{},0)
	//
	//RunCommand()

}
func RunCommand(command string) (string,string) {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var e bytes.Buffer
	cmd.Stderr =&e
	//cmd.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd.Run()
	if e.Len() != 0 && out.Len() == 0{
		return e.String(),out.String()
	}
	return "",out.String()
}

