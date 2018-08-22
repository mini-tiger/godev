package main

import (
	"github.com/gin-gonic/gin"

	"godev/mymodels/gin/funcs/user_auth"
)

func main()  {
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	route := gin.Default()
	user_auth.Load_route(route)
	route.Run(":8001")
}
