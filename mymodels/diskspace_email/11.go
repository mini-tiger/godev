package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "61566027@163.com", "一去、二三里")  // 发件人
	m.SetHeader("To",  // 收件人
		m.FormatAddress("61566027@163.com", "乔峰"),
		//m.FormatAddress("********@qq.com", "郭靖"),
	)
	m.SetHeader("Subject", "Gomail")  // 主题
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>")  // 正文

	d := gomail.NewDialer("smtp.163.com", 465, "61566027@163.com", "Taojun319")  // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}