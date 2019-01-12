package main

import (
	"godev/mymodels/ssh并发脚本/funcs"
	"godev/mymodels/ssh并发脚本/g"
	"path"
	"path/filepath"
	"runtime"
	"tjtools/logDiy"
)

func main() {
	//解析config
	_, dir, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(dir)
	logDiy.InitLog1(filepath.Join(currDir, "run.log"), 2)

	g.ParseConfig(filepath.Join(currDir, "config.json"))

	funcs.SSHRun() // 测试密码
	logDiy.Logger().Printf("密码错误的主机有:%v\n", funcs.FailHosts)
	logDiy.Logger().Printf("正确的主机有:%v\n", funcs.HostPass.Keys())
	if funcs.HostPass.IsEmpty() {
		logDiy.Logger().Fatalf("没有成功的主机")
	}

	funcs.UploadFile(filepath.Join(currDir, "1.sh"), path.Join("/tmp/", "1.sh")) // 源脚本， 目标文件,unix 使用path.join
	logDiy.Logger().Printf("上传文件错误的主机有:%v\n", funcs.FailSftpHosts)
	//ssh.Session.Stdout=os.Stdout
	//ssh.Session.Stderr=os.Stderr
	//ssh.Session.Run("touch /root/1")
	//ssh.Session.Run("ls /; ls /tmp")
	//ssh.close_session() //todo session一次运行一次run

	//terminal_run(ssh.Session)
	//ssh.close_session()
	select {}
}
