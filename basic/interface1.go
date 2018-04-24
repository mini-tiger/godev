package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	fi, err := os.Open("111")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, b, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a), b, c)
	}
}
