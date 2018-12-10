package main

import (
	"fmt"
	"net/http"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	//配置微信参数
	config := &wechat.Config{
		AppID:          "1",
		AppSecret:      "1YwqhYfsUKDoRYYJWjMRcL_T0G2ZSPCPmQ2RwhDTzcw",
		Token:          "weixin",
		EncodingAESKey: "BQptm8SueWbIj8z1NRPNSxdznzSAmRMiP54cKSmCsQh",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}
