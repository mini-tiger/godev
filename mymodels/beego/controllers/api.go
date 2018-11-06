package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

var Ns *beego.Namespace = beego.NewNamespace("/v1", //todo 第一级目录
	beego.NSCond(func(ctx *context.Context) bool { // 判断为真，才返回数据
		//fmt.Println(ctx.Input.IP())
		if ctx.Input.Domain() == "api.beego.me" || ctx.Input.IP() == "127.0.0.1" {
			return true
		}
		return false
	}),
	//beego.NSBefore(auth),
	beego.NSGet("/notallowed", func(ctx *context.Context) { // todo 第二级 /v1/notallowed
		fmt.Println(beego.AppConfig.String("testvar")) //todo 提取配置文件中的值
		ctx.Output.Body([]byte("notAllowed"))
	}),
	//beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
	//beego.NSRouter("/changepassword", &UserController{}),
	beego.NSNamespace("/shop", // todo 第二级 又分子目录  /v1/shop/d
		//beego.NSBefore(sentry),
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("this is /ship/:id"))
		}),
	),

	beego.NSNamespace("/cms", // todo 这里没有实现， /v1/cms/ 下面访问不到
		beego.NSInclude(
			//&MainController{},
			&UserController{},
		),
	),
)

//func init()  {
//	Ns :=
//		beego.NewNamespace("/v1",
//			beego.NSCond(func(ctx *context.Context) bool {
//				if ctx.Input.Domain() == "api.beego.me" {
//					return true
//				}
//				return false
//			}),
//			//beego.NSBefore(auth),
//			beego.NSGet("/notallowed", func(ctx *context.Context) {
//				ctx.Output.Body([]byte("notAllowed"))
//			}),
//			//beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
//			//beego.NSRouter("/changepassword", &UserController{}),
//			beego.NSNamespace("/shop",
//				//beego.NSBefore(sentry),
//				beego.NSGet("/:id", func(ctx *context.Context) {
//					ctx.Output.Body([]byte("notAllowed"))
//				}),
//			),
//			beego.NSNamespace("/cms",
//				beego.NSInclude(
//					&MainController{},
//					&CMSController{},
//				),
//			),
//		)
//}
