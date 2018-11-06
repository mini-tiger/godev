package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"] = append(beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: `/all/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"] = append(beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"] = append(beego.GlobalControllerRouter["godev/mymodels/beego/controllers:CC"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: `/staticblock`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godev/mymodels/beego/controllers:UserController"] = append(beego.GlobalControllerRouter["godev/mymodels/beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
