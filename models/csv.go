package main

import (
	"path/filepath"
	"os"
	log "github.com/sirupsen/logrus"
	"encoding/csv"
	"fmt"
	"io"
)

const filename = "Book1.csv"

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})
}

func main() {
	file := filepath.Join("c:\\", filename)
	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		log.Errorln(err)
	}
	readfile(f)

	f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //写入文件，追加
	if err != nil {
		log.Errorln(err)
	}
	writefile(f)
}

func readfile(f *os.File) {
	r := csv.NewReader(f)
	//fmt.Println(r.ReadAll())//全部读取 二维切片，子切片是一行
	for i := 1; ; i++ {
		r, err := r.Read()
		if err != nil || err == io.EOF {
			break
		}
		fmt.Printf("第%d行是:%v\n", i, r)
	}
}

func writefile(f *os.File) {
	w := csv.NewWriter(f)
	data := [][]string{
		{"a"},
		{"b"},
	}
	err := w.WriteAll(data)
	if err != nil {
		log.Fatalln(err)
	}
}
