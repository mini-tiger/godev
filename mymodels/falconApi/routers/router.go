// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"godev/mymodels/falconApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/alarms",
			beego.NSInclude(
				&controllers.AlarmController{},
			),
		),
		beego.NSNamespace("/single_alarms",
			beego.NSInclude(
				&controllers.SingleAlarms{},
			),
		),
	)
	//ns :=
	//	beego.NewNamespace("/v1",
	//		beego.NSNamespace("/customer",
	//			beego.NSInclude(
	//				&controllers.CustomerController{},
	//				&controllers.CustomerCookieCheckerController{},
	//			),
	//		),
	//		beego.NSNamespace("/catalog",
	//			beego.NSInclude(
	//				&controllers.CatalogController{},
	//			),
	//		),
	//	)
	beego.AddNamespace(ns)
}
