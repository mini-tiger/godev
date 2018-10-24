package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)



func main() {

	router := gin.Default()
	router.Any("/any", func(c *gin.Context) {
		//fmt.Println(c.ClientIP())
		//fmt.Println(c.Request)
		//fmt.Println(c.Request.Method)
		rStr:=fmt.Sprintf("来访IP: %s\n",c.ClientIP())
		rStr+=fmt.Sprintf("主机名:%s\n远程IP:%s\n",c.Request.Host,c.Request.RemoteAddr)
		rStr+=fmt.Sprintf("请求的URL: %s\n",c.Request.URL)
		rStr+=fmt.Sprintf("使用的浏览器客户端:%s\n",c.Request.UserAgent())
		rStr+=fmt.Sprintf("请求方法: %s\n",c.Request.Method)
		for k,v:=range c.Request.Header{
			rStr+=fmt.Sprintf("请求头:%s ,value:%v\n",k,v)
		}
		//log.Println("Done! in path" + c.Request.URL.Path)
		switch c.Request.Method {
		case "GET":
			c.String(http.StatusOK,rStr)
			return
		case "POST":
			c.String(http.StatusOK,rStr)
			return

		}
		//c.String(http.StatusOK,"this is sync")
	})



	router.Run(":8001")
}
