package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

// CMS API
type CMSController struct {
	beego.Controller
}

type CC struct { // 派生一个子类，继承cmscontroller 的Prepare，finish方法
	CMSController
}


func (c *CC) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
	c.Mapping("get", c.Get)
}

func (this *CMSController)Prepare()  {
	fmt.Println("this is prepare ")
}

func (this *CMSController)Finish()  {
	fmt.Println("this is finish ")
}

// @router /get [get]
func (this *CC)Get()  {
	fmt.Println(3)
	this.Ctx.WriteString(fmt.Sprintf("input is Get"))
}
// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock [get]
func (this *CC) StaticBlock() {
	fmt.Println(1)
	this.Ctx.WriteString(fmt.Sprintf("input is "))
}

// @router /all/:key [get]
func (this *CC) AllBlock() {
	fmt.Println(2)
	s:=this.Ctx.Input.Param(":key")
	this.Ctx.WriteString(fmt.Sprintf("input is %s",s))
}