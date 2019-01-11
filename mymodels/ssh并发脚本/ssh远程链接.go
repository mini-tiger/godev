package main

import (
	"fmt"
	"godev/mymodels/ssh并发脚本/funcs"
	"godev/mymodels/ssh并发脚本/g"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"tjtools/logDiy"
)

func main() {
	//解析config
	_, dir, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(dir)
	logDiy.InitLog1(filepath.Join(currDir, "run.log"), 2)

	g.ParseConfig(filepath.Join(currDir, "config.json"))

	funcs.SSHRun() // 测试密码
	fmt.Printf("密码错误的主机有:%v\n", funcs.FailHosts)
	fmt.Printf("正确的主机有:%v\n", funcs.HostPass.Keys())

	funcs.UploadFile() // 源脚本， 目标文件

	//ssh.Session.Stdout=os.Stdout
	//ssh.Session.Stderr=os.Stderr
	//ssh.Session.Run("touch /root/1")
	//ssh.Session.Run("ls /; ls /tmp")
	//ssh.close_session() //todo session一次运行一次run

	//terminal_run(ssh.Session)
	//ssh.close_session()
	select {}
}

func terminal_run(session *ssh.Session) {
	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}
	e, err := session.StderrPipe()
	if err != nil {
		panic(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal 建立终端
	if err := session.RequestPty("vt100", 40, 80, modes); err != nil { //term:xerm 是彩色显示
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	in, out := MuxShell(w, r, e)
	if err := session.Shell(); err != nil { //打开仿真shell
		log.Fatal(err)
	}
	//<-out 通信out第一次返回的是 linux 登录信息,可以跳过
	in <- "ls /"
	in <- "ls /tmp"

	in <- "exit"

	for {
		if k, ok := <-out; ok {
			fmt.Printf("%s\n", k) //todo 所有out通道中记录的 返回信息 打开出来
		} else {
			break
		}
	}

	session.Wait()

}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 0)
	out := make(chan string, 4)
	var wg sync.WaitGroup
	wg.Add(1) //shell 退出前，Shell的进程
	go func() {
		for cmd := range in { //todo in通道中 所有 需要执行的命令 依次执行
			wg.Add(1)
			w.Write([]byte(cmd + "\n")) //w 是管道输入
			wg.Wait()                   //等待命令完成
		}
	}()

	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:]) //todo 标准输出管道的 内容，stdoutpipe,是io.Reader接口有reader方法，将传入的[]byte 写入
			if err != nil {
				if err == io.EOF { //如果EOF 退出
					fmt.Println("exit")
					//wg.Done()
				}
				//fmt.Println(err.Error())
				close(in)
				close(out)
				return
			}
			t += n //每次命令结果 追加至buf
			result := string(buf[:t])
			if strings.Contains(result, "password:") ||
				strings.Contains(result, "#") { //匹配是否执行完成
				out <- result
				t = 0 //t是临时存 当前命令返回的结果，清空
				wg.Done()
			}
		}
	}()
	return in, out
}
