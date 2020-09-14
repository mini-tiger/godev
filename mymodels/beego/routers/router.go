// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.

package routers

import (
	"godev/mymodels/beego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	// nametuple 用在API路由
	beego.AddNamespace(controllers.Ns)

    beego.Router("/", &controllers.MainController{})

    //固定路由 http://127.0.0.1:8081/study
	beego.Router("/study", &controllers.Study{}) // 默认对应 Get POST 等方法

	//自定义方法路由
	//beego.Router("/api/get",&controllers.Study1{},"*:AllFunc;post:PostFunc")  // 不同方法对应不同函数
	beego.Router("/api/get",&controllers.Study1{},"*:GetAbc") //所有方法同一函数
	//beego.Router("/api/get",&controllers.Study1{},"get,post:getabc") //多个方法同一函数

	// 自动匹配, strudy2 绑定Login方法  会有/study2/login路由
	beego.AutoRouter(&controllers.Study2{})

    // todo 注释路由，在controller 文件中
    //  dev 模式下会进行生成，生成的路由放在 “/routers/commentsRouter.go” 文件中
    beego.Include(&controllers.CC{})




	// 以下是示例，直接返回
	beego.Get("/abc",func(ctx *context.Context){ //基本GET方法
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Any("/foo",func(ctx *context.Context){ // 响应所有方法
		//ctx.Output.Body([]byte("bar"))
		ctx.WriteString("hello ")
	})


}
