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
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	filename = "_tmp.txt"
)

//todo  win环境运行前注释 func linux_freesize

type pan string

var pf map[pan]uint64
var os_type string = runtime.GOOS

func win_run() {
	pf = make(map[pan]uint64, 0)
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
		pf[pan(temp_slice[1])] = freesize

	}

}

func loop_write(f *os.File, freesize uint64) {
	for i := uint64(0); i < freesize; i++ {
		//ioutil.WriteFile(file,[]byte{65},0644)
		f.Write([]byte{65})
		time.Sleep(1 * time.Second)
	}
}

func create_file_sub(file string, freesize uint64) {
	log.Infof("write file path:%s", file)
	//create file

	f, err1 := os.OpenFile(file, os.O_CREATE|os.O_SYNC|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 != nil {
		log.Errorf("file err:%s\n", err1)
	}
	loop_write(f, freesize)
	//}

	//}

}

func create_file() {
	for pan, freesize := range pf {
		if strings.Contains(os_type, "linux") {
			file := filepath.Join(string(pan), filename)
			//fmt.Println(file)
			go create_file_sub(file, freesize)
		} else {
			file := filepath.Join(string(pan), "\\", filename)
			go create_file_sub(file, freesize)
		}

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
	if strings.Contains(os_type, "win") {
		win_run()
		create_file()
	}
	if strings.Contains(os_type, "linux") {
		linux_run()
		create_file()
	}
	select {}
}

func linux_run() {
	pf = make(map[pan]uint64, 0)
	data, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		if err != io.EOF {
			log.Errorf("linux /proc/mounts err:%s\n", err)
		}

	}
	bb := bytes.NewBuffer(data)
	bf := bufio.NewReader(bb)
	for {
		line, err := bf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Errorf("linux /proc/mounts err:%s\n", err)
			}
			break
		}
		tmp_s := strings.Fields(line)
		if strings.Contains(tmp_s[0], "/dev") {
			//fmt.Println(tmp_s)

			//pf[pan(tmp_s[1])] = linux_freesize(tmp_s[1])
		}
	}

}

//func linux_freesize(pan string) (freeze uint64) {
//
//	fs := syscall.Statfs_t{}
//	err := syscall.Statfs(pan, &fs)
//	if err != nil {
//		return
//	}
//	All := fs.Blocks * uint64(fs.Bsize)
//	Free := fs.Bfree * uint64(fs.Bsize)
//	Used := All - Free //unit   byte
//	//fmt.Println(Used)
//	freeze = Used
//	return
//
//}
