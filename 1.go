package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
func main() {
	str := "12345678"
	//md5Str := md5V(str)
	//fmt.Println(md5Str)
	//fmt.Println(md5V2(str))
	fmt.Println(md5V3(str))

	s := "test8|dashboard"
	b := []byte(s)

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("enc=[%s]\n", sEnc)

	sDec, err := base64.StdEncoding.DecodeString("dGVzdDZ8ZGFzaGJvYXJk")
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	} else {
		fmt.Printf("dec=[%s]\n", sDec)
	}

}
