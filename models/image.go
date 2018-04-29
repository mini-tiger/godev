package main

import (
	"./thumbnail"
	// "bufio"
	"fmt"
	// "log"
	// "os"
	// "time"
)

/*
C:\godev\models>go run image.go
img.jpg
img.thumb.jpg

*/

var Jpg thumbnail.Jpginfo

type item struct {
	thumbfile string
	err       error
}

// func process(j thumbnail.Jpginfo) {

// 	// thumb, err := thumbnail.ImageFile(j)
// 	// if err != nil {
// 	// 	log.Print(err)
// 	var it item
// 	it.thumbfile, it.err = thumbnail.ImageFile(j)
// 	fsch <- it
// }

func main() {
	// input := bufio.NewScanner(os.Stdin)
	fsch := make(chan item)
	fs := []string{"C:/godev/models/pic/img.jpg", "C:/godev/models/pic/img1.jpg", "C:/godev/models/pic/img2.jpg"}

	for _, f := range fs {
		// Jpg.Filename = input.Text
		Jpg.Filename = f
		Jpg.Hight = 128
		Jpg.Weight = 128
		go func(j thumbnail.Jpginfo) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(j)
			fsch <- it
		}(Jpg)
	}
	// if err := input.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	for range fs {
		it := <-fsch
		fmt.Println(it.thumbfile, it.err)
	}
}
