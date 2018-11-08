package user_auth

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func Post(c *gin.Context)  {
	inputs := APILoginInput{}
	if err := c.Bind(&inputs); err != nil {
		//h.JSONR(c, badstatus, "name or password is blank") 可以写成统一返回错误的方法
		c.String(http.StatusOK, "Hello World")
	}
	name := inputs.Username
	password := inputs.Passwd
	fmt.Println(name,password)
	//通过DB验证，return json
	//这里直接resturn
	if inputs.Username == ""{
		fmt.Println(111111111)
		c.String(http.StatusBadRequest, "Hello World")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   inputs.Username,
		"passwd": inputs.Passwd,
	})
	return
}

func Login(c *gin.Context)  {
	contentType := c.Request.Header.Get("Content-Type") //获取头文件
	panic(123)
	fmt.Println(contentType)
	inputs := APILoginInput{}
	if err := c.Bind(&inputs); err != nil {
		//h.JSONR(c, badstatus, "name or password is blank") 可以写成统一返回错误的方法
		c.String(http.StatusOK, "Hello World")
		return
	}
	name := inputs.Username
	password := inputs.Passwd
	fmt.Println(name,password)
	//通过DB验证，return json
	//这里直接resturn
	c.JSON(http.StatusOK, gin.H{
		"user":   inputs.Username,
		"passwd": inputs.Passwd,
	})
	return
}