package main

import (
	"os/exec"
	"fmt"
	"time"
	"reflect"
	"bytes"
)

func main()  {
	run()
}

func run()  {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd %s;./%s", "/home", "tomcat_start.sh"))

	go func(cmd *exec.Cmd) { // 实时监控判断脚本运行状态，打印日志
		for i := 1; ; i++ {
			var d int = 10
			time.Sleep(time.Duration(d) * time.Second)
			c := cmd.Process
			v := reflect.ValueOf(c).Elem()
			vn := v.FieldByName("isdone")
			if (vn.Uint() == uint64(1)) {
				fmt.Printf("脚本运行完成%s,运行参数:%+v\n", cmd.Args, c)
				break
			} else {
				fmt.Printf("正在运行的脚本%s,已运行%d秒,运行参数:%+v\n", cmd.Args, i*d, c)
			}
		}
	}(cmd)

	var out bytes.Buffer
	cmd.Stdout = &out

	var e bytes.Buffer
	cmd.Stderr = &e
	//cmd.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd.Run()
	fmt.Println(out.String())
	fmt.Println(e.String())
}