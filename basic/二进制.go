package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"path/filepath"
)

type Register struct {
	ACTION int32
	SID    int32
}

var M []interface{}

var (
	BaseDir string
)

func ExampleWrite() []byte {
	buf := new(bytes.Buffer)

	var info Register
	info.ACTION = 20004
	info.SID = 6

	err := binary.Write(buf, binary.LittleEndian, info)
	if err != nil {
		fmt.Println("binary.Write failed：", err)
	}
	fmt.Printf("% x\n", buf.Bytes())
	//写入文件
	ioutil.WriteFile(filepath.Join(BaseDir, "bin_file.txt"), buf.Bytes(), 0755)
	return buf.Bytes()
}

func ExampleRead(b []byte) {
	var info Register
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, &info)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(info)
	fmt.Println()

	// 从文件读取
	ds, err := os.Open(filepath.Join(BaseDir, "bin_file.txt"))
	if err != nil {
		fmt.Errorf("%s\n", err)
	}

	err1 := binary.Read(ds, binary.LittleEndian, &info)
	if err1 != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(info)
}

func main() {
	_, ParsDir, _, _ := runtime.Caller(0)
	BaseDir = filepath.Dir(ParsDir)
	buf := ExampleWrite()
	ExampleRead(buf)
	fmt.Println()
	fmt.Println("==========================slice read write============================")
	buf = SliceWrite()
	SliceRead(buf)
}

func SliceRead(b []byte) {
	buf := bytes.NewBuffer(b)
	i:= make([]int8,3) // todo 要预先定义好数组个数 或者 var i [3]int8
	err := binary.Read(buf, binary.LittleEndian, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(i)
	// 从文件读取
	ds, err := os.Open(filepath.Join(BaseDir, "bin_slice_file.txt"))
	if err != nil {
		fmt.Errorf("%s\n", err)
	}

	err1 := binary.Read(ds, binary.LittleEndian, &i)
	if err1 != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(i)
}
func SliceWrite() []byte {
	buf := new(bytes.Buffer)

	d := []int8{1, 2, 3} // todo 要预先定义好写入的元素个数
	for _, v := range d {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("% x\n", buf.Bytes())

	//写入文件
	ioutil.WriteFile(filepath.Join(BaseDir, "bin_slice_file.txt"), buf.Bytes(), 0755)
	return buf.Bytes()
}
