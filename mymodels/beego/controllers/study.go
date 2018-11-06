package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

type Study struct {
	beego.Controller
}

type User struct {
	Name  interface{} `form:"username",json:"name"`
	Age   int `form:"age",json:"age"`
}

func (c *Study) Get() {
	fmt.Println(c.GetString("abcpost")) // todo 获取 post, get参数
	fmt.Println(c.Input().Get("abcpost"))
	c.Data["data1"] = "taojun"
	c.Data["data2"] = "this is test data"
	c.TplName = "study/index.tpl" //
}

func (c *Study) Post() {
	//u := User{}
	//if err := c.ParseForm(&u); err != nil { // todo formdata
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
	fmt.Println(string(c.Ctx.Input.RequestBody))
	u:=User{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	fmt.Println(u)

	//c.Data["data1"] = "taojun"
	//c.Data["data2"] = "this is test data"
	//c.TplName = "study/index.tpl" //
	c.Data["json"]=u
	c.ServeJSON()

}