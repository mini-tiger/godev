package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"reflect"
)

func main() {
	f, _ := os.OpenFile("/home/go/GoDevEach/works/haifei/syncHtmlYWReport/HisData/bak/2021-02-12/HFserverdxc01_yuebao_7DAY_1COPY_Chinese.html", os.O_RDONLY, 0755)

	defer func() {
		_ = f.Close()
	}()

	dom, _ := goquery.NewDocumentFromReader(f)
	fmt.Printf("%p\n", dom)
	dom.Empty()
	fmt.Printf("%p\n", dom)
	Clear(dom)
	fmt.Printf("%p\n", dom)
}

func Clear(v interface{}) { // 必须传入指针
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}
