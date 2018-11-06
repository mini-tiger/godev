package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	beego.Controller
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	fmt.Println(123)
	u.Data["json"] = "logout success"
	u.ServeJSON()
}