package user_auth

import (
	"github.com/gin-gonic/gin"
	"godev/mymodels/gin/funcs/utils"
	"log"
)

func Routes(g *gin.Engine) {
	//u := g.Group("/api/v1/user", gin.BasicAuth(gin.Accounts{ // 所有此组下的 接口 要认证，herader加入 认证信息 Authorization: Basic Zm9vOmJhcg==  ，base64编码
	//	"foo":    "bar", //用户名：密码 https://blog.csdn.net/puss0/article/details/81003400
	//	"austin": "1234",
	//	"lena":   "hello2",
	//	"manu":   "4321",
	//})) //建立URL组
	u := g.Group("/api/v1/user") //建立URL组
	//u.GET("/auth_session", AuthSession)
	u.Use(utils.Test_midddle) //单独组加中间件,中间件与视图不要放在一个包中
	//u.Use(gin.RecoveryWithWriter(os.Stdout))
	//u.Any("/login123", RecoverPanic(Post)) //组下面URL 都是 组的URL开头 POST方法
	u.POST("/login123", RecoverPanic(Post)) //组下面URL 都是 组的URL开头 POST方法
	//fmt.Printf("%T,%+v\n",a,a)
	u.Handle("GET", "/login", RecoverPanic(Login))

	//u.GET("/logout", Logout)
}

// todo 外层抓错误 返回
func RecoverPanic(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[ERROR][%v] url:%s, panic:%v\n", context.Request.RemoteAddr, context.Request.URL, err)
			}
		}()
		handle(context)
	}
}
