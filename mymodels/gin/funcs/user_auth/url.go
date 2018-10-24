package user_auth

import (
	"github.com/gin-gonic/gin"
	"godev/mymodels/gin/funcs/utils"
)

func Routes(g *gin.Engine)  {
	u := g.Group("/api/v1/user") //建立URL组
	//u.GET("/auth_session", AuthSession)
	u.Use(utils.Test_midddle)//单独组加中间件,中间件与视图不要放在一个包中
	u.POST("/login", Login) //组下面URL 都是 组的URL开头 POST方法

	//u.GET("/logout", Logout)

}
