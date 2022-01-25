package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decoded, err := base64.StdEncoding.DecodeString("dGFvLmp1bjpUYW9qdW4yMDc=")
	decodestr := string(decoded)
	fmt.Println(decodestr, err)
}
