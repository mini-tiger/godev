package controllers

import (
	"github.com/astaxie/beego"
	"godev/mymodels/falconApi/funcs"
	"log"
)

// Operations about object
type AlarmController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	count   formData   int false       "count"
// @Success 200 {object} funcs.CountTop
// @Failure 403 body is empty
// @router / [post]
func (o *AlarmController) Post() {
	//fmt.Println(o.Ctx.Input.Param(":count"))
	count, err := o.GetInt("count")
	if err != nil {
		log.Println(err)
	}
	m := funcs.GetEventCase(count)
	o.Data["json"] = m
	//var ob models.Object
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//objectid := models.AddOne(ob)
	//o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
//func (o *AlarmController) Get() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	if objectId != "" {
//		ob, err := models.GetOne(objectId)
//		if err != nil {
//			o.Data["json"] = err.Error()
//		} else {
//			o.Data["json"] = ob
//		}
//	}
//	o.ServeJSON()
//}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
//func (o *AlarmController) GetAll() {
//	obs := models.GetAll()
//	o.Data["json"] = obs
//	o.ServeJSON()
//}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
//func (o *ObjectController) Put() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	var ob models.Object
//	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
//
//	err := models.Update(objectId, ob.Score)
//	if err != nil {
//		o.Data["json"] = err.Error()
//	} else {
//		o.Data["json"] = "update success!"
//	}
//	o.ServeJSON()
//}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
//func (o *ObjectController) Delete() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	models.Delete(objectId)
//	o.Data["json"] = "delete success!"
//	o.ServeJSON()
//}

type SingleAlarms struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	endpoint   formData   string true       "endpoint"
// @Param	pageno   formData   int true       "第几页"
// @Param	pagesize   formData   int true       "每页多少个"
// @Success 200 {object} funcs.Page
// @Failure 403 endpoint is empty
// @router / [post]
func (o *SingleAlarms) Post() {
	//fmt.Println(o.Ctx.Input.Param(":count"))
	ep := o.GetString("endpoint")
	pn, err := o.GetInt("pageno")
	if err != nil {
		var e funcs.Page
		PageReError(err, o, e)
		log.Printf("获取 PageNo err:%s", err)
	}
	ps, err := o.GetInt("pagesize")
	if err != nil {
		var e funcs.Page
		PageReError(err, o, e)
		log.Printf("获取 PageSize err:%s", err)
		return
	}

	m := funcs.GetSingleEvent(ep, pn, ps)
	log.Printf("return data %+v\n", m)
	o.Data["json"] = m
	o.ServeJSON()
}

func PageReError(e error, o *SingleAlarms, p funcs.Page) {
	p.Error = e.Error()
	o.Data["json"] = p
	o.ServeJSON()

}
