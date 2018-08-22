package user_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"godev/mymodels/gin/funcs/utils"
)

func Load_route(g *gin.Engine)  {
	// 静态资源返回
	g.StaticFile("favicon.ico", "./resources/favicon.ico")
	// 静态资源目录
	g.Static("/css", "/var/www/css")

	g.Use(utils.CORS()) //加载中间件
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm Falcon+ (?A?)")
	})
	//加载用户路径 与中间件
	Routes(g)
}
