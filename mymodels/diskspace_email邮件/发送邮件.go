package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"mime"
)

func main() {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "61566027@163.com", "一去、二三里")  // 发件人
	msg.SetHeader("To",  // 收件人
		msg.FormatAddress("61566027@163.com", "乔峰"),
		//m.FormatAddress("********@qq.com", "郭靖"),
	)
	msg.SetHeader("Subject", "Gomail")  // 主题
	msg.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>")  // 正文

	//添加附件
	name := "中文.xlsx"
	msg.Attach("C:\\work\\go-dev\\src\\godev\\mymodels\\1.xlsx",
		gomail.Rename(name),
		gomail.SetHeader(map[string][]string{
			"Content-Disposition": []string{ //中文文件名
				fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", name)),
			},
		}),
	)

	d := gomail.NewDialer("smtp.163.com", 465, "61566027@163.com", "IORQCTROCDAKWPZS")  // 发送邮件服务器、端口、发件人账号、授权码
	if err := d.DialAndSend(msg); err != nil {
		panic(err)
	}


}