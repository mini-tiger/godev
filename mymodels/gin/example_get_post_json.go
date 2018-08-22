package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"github.com/gin-gonic/gin/binding"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") //http://127.0.0.1:8001/user/ddd
		//Hello ddd
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") //获取get 方法参数，并给出默认值
		lastname := c.Query("lastname")//只有获取，默认空字符串

//http://127.0.0.1:8001/welcome?firstname=中国&lastname=中国
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message") //post方式 获取
		nick := c.DefaultPostForm("nick", "anonymous") //post方式 获取，有默认值

		//JSON 返回json数据，gin.H 封装json数据包
		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type") //获取头文件
		fmt.Println(contentType)
		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err = c.BindWith(&user, binding.Form)
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})

	})

	router.Run(":8001")
}
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}