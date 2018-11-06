package controllers

import (
	"github.com/astaxie/beego"
)

type Study1 struct {
	beego.Controller
}

func (c *Study1) GetAbc() {
	c.Data["data1"] = "taojun"
	c.Data["data2"] = "this is test data"
	//c.TplName = "study/index.tpl" // todo 目录名文件名小写
	// todo 不指定TplName 会自动去 模板目录的 Controller/<方法名>.tpl 查找 study1/getabc.tpl
}
