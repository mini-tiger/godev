package main

import (
	"github.com/gin-gonic/gin"
	"godev/mymodels/gin/funcs/user_auth"

	"github.com/betacraft/yaag/yaag"
	yaag_gin "github.com/betacraft/yaag/gin"

)

func main()  {
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	route := gin.Default()

	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Gin",
		DocPath:  "apidoc1.html",
		BaseUrls: map[string]string{"Production": "", "Staging": "http://127.0.0.1:8001"},
	})
	route.Use(yaag_gin.Document())

	user_auth.Load_route(route)
	route.Run(":8001")
}
