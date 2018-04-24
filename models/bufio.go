package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "path/filepath"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	fmt.Println(111)
	return len(p), nil
}

//!-bytecounter

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!\n")
	b.WriteTo(os.Stdout)

	contents, err := ioutil.ReadFile("111")
	fmt.Println(strings.Count(string(contents), "\n")) //一共几行
	fi, err := os.Open("111")
	// fmt.Println(*fi)
	// c, _ := ioutil.ReadAll(fi)
	// fmt.Println(string(c))

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	bs := bufio.NewScanner(br)
	bs.Split(bufio.ScanWords)
	var rw int

	for bs.Scan() {
		rw++
		// fmt.Println(bs.Text())
		// if strings.Count(bs.Text(), "\n") > 0 {
		// 	rn++
		// }

	}
	fmt.Println(rw) //一共几个单词
}
