package bussiness

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
	"reflect"
	"godev/mymodels/mongo_sync/models"
)

type CC_APPBase models.CC_ApplicationBase

const collectionNameApp  = "cc_ApplicationBase"

type CCApp struct {
	collect          string
	src, dst         *mgo.Collection
	addSlice         []*CC_APPBase
	delSlice         []*CC_APPBase
	updateSlice      []*CC_APPBase
	table            CC_APPBase
	srcdata, dstdata []*CC_APPBase
	srcMap, dstMap   map[bson.ObjectId]*CC_APPBase
}

func (self *CCApp) CollectionName() string{
	return self.collect
}
func (self *CCApp) Init(sessionSrc, sessionDst *mgo.Session) {
	self.collect = collectionNameApp
	self.src = sessionSrc.DB("cmdb").C(self.collect)
	self.dst = sessionDst.DB("cmdb1").C(self.collect)
	self.srcdata = make([]*CC_APPBase, 0)
	self.dstdata = make([]*CC_APPBase, 0)
	self.addSlice = make([]*CC_APPBase, 0)
	self.delSlice = make([]*CC_APPBase, 0)
	self.updateSlice = make([]*CC_APPBase, 0)
	self.srcMap = make(map[bson.ObjectId]*CC_APPBase, 0) // map[id]person{name,phone}   person不包含id, id当作KEY
	self.dstMap = make(map[bson.ObjectId]*CC_APPBase, 0)
}

func (self *CCApp) DataAll() bool {

	self.src.Find(nil).All(&self.srcdata)
	//c.Find(nil).Select(bson.M{"_id":1}).All(&users)
	if len(self.srcdata) == 0 {
		return false
	}
	self.dst.Find(nil).All(&self.dstdata)
	return true
}

func (self *CCApp) DiffSlice() bool {
	if reflect.DeepEqual(self.srcdata, self.dstdata) {
		return false //不需要往下同步数据
	}

	for _, v := range self.dstdata {
		//p := &self.table
		//p.Name = v.Name
		//p.Phone = v.Phone
		self.dstMap[v.Id_] = v
	}

	for _, v := range self.srcdata {
		//p := &self.table
		//p.Name = v.Name
		//p.Phone = v.Phone
		self.srcMap[v.Id_] = v
	}

	for _, v := range self.srcdata {
		if _, ok := self.dstMap[v.Id_]; ok { //源数据id在目标数据中有，应该比对其它相关字段
			tmp := self.dstMap[v.Id_]
			if v.Bk_biz_developer == tmp.Bk_biz_developer && v.Bk_biz_id == tmp.Bk_biz_id &&
				v.Bk_biz_maintainer == tmp.Bk_biz_maintainer && v.Bk_biz_name == tmp.Bk_biz_name &&
				v.Bk_biz_productor == tmp.Bk_biz_productor && v.Bk_biz_tester == tmp.Bk_biz_tester &&
				v.Bk_supplier_account == tmp.Bk_supplier_account && v.Bk_supplier_id == tmp.Bk_supplier_id &&
				v.Create_time == tmp.Create_time && v.Defaultt == tmp.Defaultt &&
				v.Language == tmp.Language && v.Last_time == tmp.Last_time &&
				v.Life_cycle == tmp.Life_cycle && v.Operator == tmp.Operator && v.Time_zone == tmp.Time_zone { //其它字段都符合则跳过
				continue
			} else {
				//fmt.Println(self.detMap[v.Id_].Name==v.Name)
				//fmt.Println(v.Phone == self.dstMap[v.Id_].Phone)
				//fmt.Println(self.dstMap[v.Id_])
				//fmt.Println(v)
				self.updateSlice = append(self.updateSlice, v)
			}
		} else {

			self.addSlice = append(self.addSlice, v) //源数据有,目标没有
		}
	}

	for _, vv := range self.dstdata {
		if _, ok := self.srcMap[vv.Id_]; ok { //目标有，源数据中有,不考虑目标库 同步到 源库
			continue
		} else {
			self.delSlice = append(self.delSlice, vv) //目标有,源没有，应该从目标 删除
			//continue
		}
	}
	return true
}

func (self *CCApp) Run() {

	//c1.Find(nil).Select(bson.M{"_id":1}).All(&users1)

	//fmt.Println(users[0].Id_)

	//if len(users1) == len(users){ //判断长度相等意义 不大
	//	fmt.Println("相等")
	//}
	////a:=[]int{1}
	////a1:=[]int{1}
	if len(self.srcdata) == 0 {
		return //源数据无，直接返回
	}

	//fmt.Println(self.addSlice)
	//fmt.Println(self.delSlice)
	//fmt.Println(self.updateSlice)
	for _, v := range self.addSlice {
		//p := Person{v.Id_, v.Name, v.Phone}
		Log.Info("insert data %s",v.Id_)
		self.dst.Insert(&v)

	}
	for _, v := range self.delSlice {
		//p := Person{v.Id_, v.Name, v.Phone}
		Log.Info("remove data %s",v.Id_)
		self.dst.Remove(bson.M{"_id": v.Id_})

	}

	for _, v := range self.updateSlice {
		//p := Person{v.Id_, v.Name, v.Phone}
		Log.Info("update data %s",v.Id_)
		err := self.dst.Update(bson.M{"_id": v.Id_}, v)
		if err != nil {
			fmt.Println(err)
		}


	}
}
