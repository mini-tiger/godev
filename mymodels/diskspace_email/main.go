package main

import (
	"bufio"
	"bytes"
	_ "gitee.com/taojun319/godaemon"
	log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"fmt"
	"time"
	"gopkg.in/gomail.v2"
	_ "gitee.com/taojun319/godaemon"
)

//todo  win环境运行前注释 func linux_freesize

const (
	warn float64 = 20.0
)

type pan string

var pf map[pan]string

//linux
var upf map[string]float64 = make(map[string]float64,0)
var os_type string = runtime.GOOS

type p struct {
	Pf map[string]float64
}

var cp chan *p = make(chan *p, 0)


func win_run() {
	pf = make(map[pan]string, 0)
	cmd := exec.Command("cmd", "/c", "wmic LOGICALDISK get name,size,freespace")
	o, err := cmd.Output()
	if err != nil {
		log.Printf("err:%s\n", err)
	}
	out := bytes.NewBuffer(o)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Errorf("err:%s\n", err)
			break
		}
		if strings.Contains(line, "Free") {
			continue
		}

		temp_slice := strings.Fields(line)
		//fmt.Println(temp_slice)
		if len(temp_slice) != 3 {
			continue
		}
		freesize, err := strconv.ParseUint(temp_slice[0], 10, 64)
		if err != nil {
			log.Errorf("err:%s\n", err)
		}
		pf[pan(temp_slice[1])] = string(freesize)

	}

}

func main() {

	//fmt.Println(os_type)
	//if deamon {
	f, _ := os.OpenFile("./log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	//ff,_:=f.Stat()
	//
	//fmt.Println(ff.Name())
	log.SetOutput(f)
	//}else{
	//	log.SetOutput(os.Stdout)
	//}



	log.SetLevel(log.DebugLevel) //级别
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})
	go func() {
		for {
			if strings.Contains(os_type, "win") {
				win_run()

			}
			if strings.Contains(os_type, "linux") {
				go email_run()
				linux_run()

			}
		}
	}()
	select {}
}

func email_run()  {
	for  {
		a:=<-cp
		fmt.Println(a)
		//msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		//	"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
		for k,v :=range a.Pf {
			SendToMail("smtp.163.com", "Taojun319", "61566027@163.com", "61566027@163.com", 465, k,v)
		}
		}
}

func SendToMail(host,passwd,from,to string ,port int,p string,l float64) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "61566027@163.com", "")  // 发件人
	m.SetHeader("To",  // 收件人
		m.FormatAddress("61566027@163.com", ""),
		//m.FormatAddress("********@qq.com", "郭靖"),
	)
	m.SetHeader("Subject", "磁盘报警")  // 主题
	tmp:=fmt.Sprintf("%s 使用率 %2f",p,l)
	m.SetBody("text/html", fmt.Sprintf("<h2>%s </h2>",tmp))  // 正文

	d := gomail.NewDialer(host,port,from,passwd)  // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		return(err)
	}
	return nil
}


func linux_run() {

	data, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		if err != io.EOF {
			log.Errorf("linux /proc/mounts err:%s\n", err)
		}

	}

	bb := bytes.NewBuffer(data)
	bf := bufio.NewReader(bb)
	for {
		upf = make(map[string]float64,0)
		line, err := bf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Errorf("linux /proc/mounts err:%s\n", err)
			}
			break
		}
		tmp_s := strings.Fields(line)

		if strings.Contains(tmp_s[0], "/dev") {

			s := linux_freesize(tmp_s[1])

			if f, err := strconv.ParseFloat(s, 64); err == nil { //转换类型
				//fmt.Printf("%T,%v\n", f, f)
				if f > warn { //判断 是否超过阈值
					upf[tmp_s[1]] = f
				}
			}

		}
		if len(upf)>0{
			pp:=&p{}
			pp.Pf=upf
			cp <- pp
		}

	}
	time.Sleep(2*time.Second)
}

func linux_freesize(pan string) (string) {

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(pan, &fs)
	if err != nil {
		return ""
	}
	All := fs.Blocks * uint64(fs.Bsize)
	Free := fs.Bfree * uint64(fs.Bsize)
	Used := All - Free //unit   byte
	//fmt.Println(Used)
	//fmt.Println(All)
	//freeze = Used

	return fmt.Sprintf("%2f", float64(Used)/float64(All)*100)

}
