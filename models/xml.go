package main

import (
	"encoding/xml"
	"fmt"
	"io"
	// "io/ioutil"
	// "bytes"
	"os"
	"strings"
)

// var x string

// x=`<Person><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`

func main() {
	/*
		C:\godev\models>go run xml.go div div h2
		html body div div div h2: XML 基础
		html body div div div h2: XML JavaScript
		html body div div div h2: XML 高级
		html body div div div h2: XML 实例/测验
		html body div div div h2: 建站手册
		html body div div div h2 a: 关于 W3School
		html body div div div h2 a: 帮助 W3School
	*/
	openfile() //C:\godev\models>go run xml.go
	// read_string()
}

func read_string() {
	var t xml.Token
	var err error

	input := `<Person><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	inputReader := strings.NewReader(input)

	// 从文件读取，如可以如下：
	// content, err := ioutil.ReadFile("s.xml")
	// decoder := xml.NewDecoder(bytes.NewBuffer(content))

	decoder := xml.NewDecoder(inputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}
	}
}

func openfile() {
	// xx := strings.NewReader(x)
	p := []string{"div", "div", "h2"}
	xx, err := os.Open("xx.xml")
	if err != nil {
		fmt.Println(err)
		// return false
	}
	dec := xml.NewDecoder(xx)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement: // 例如<div> API保证与EndElement 是一对
			stack = append(stack, tok.Name.Local) // push
			for _, attr := range tok.Attr {       //Attr代表一个XML元素的一条属性（Name=Value）
				attrName := attr.Name.Local

				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement: // 例如</div>
			stack = stack[:len(stack)-1] // pop
		case xml.CharData: //不作为xml的节点输出,把该字段对应的值作为字符输出,CharData类型代表XML字符数据（原始文本）
			if containsAll(stack, p) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
			// case xml.Comment: //,comment” 输出xml中的注释
			// 	fmt.Printf("%s\n", tok)

		}

	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
