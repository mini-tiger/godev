package bussiness

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
	"reflect"
	"godev/mymodels/mongo_sync/models"
)

type CC_H models.CC_History

var collectionNameHistory  = "cc_History"

type CCHistory struct {
	collect          string
	src, dst         *mgo.Collection
	addSlice         []*CC_H
	delSlice         []*CC_H
	updateSlice      []*CC_H
	table            CC_H
	srcdata, dstdata []*CC_H
	srcMap, dstMap   map[bson.ObjectId]*CC_H
}

func (self *CCHistory) CollectionName() string{
	return self.collect
}
func (self *CCHistory) Init(sessionSrc, sessionDst *mgo.Session) {
	self.collect = collectionNameHistory
	self.src = sessionSrc.DB("cmdb").C(self.collect)
	self.dst = sessionDst.DB("cmdb1").C(self.collect)
	self.srcdata = make([]*CC_H, 0)
	self.dstdata = make([]*CC_H, 0)
	self.addSlice = make([]*CC_H, 0)
	self.delSlice = make([]*CC_H, 0)
	self.updateSlice = make([]*CC_H, 0)
	self.srcMap = make(map[bson.ObjectId]*CC_H, 0) // map[id]person{name,phone}   person不包含id, id当作KEY
	self.dstMap = make(map[bson.ObjectId]*CC_H, 0)
}

func (self *CCHistory) DataAll() bool {

	self.src.Find(nil).All(&self.srcdata)
	//c.Find(nil).Select(bson.M{"_id":1}).All(&users)
	if len(self.srcdata) == 0 {
		return false
	}
	self.dst.Find(nil).All(&self.dstdata)
	return true
}

func (self *CCHistory) DiffSlice() bool {
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
			if v.Id == tmp.Id && v.Create_time == tmp.Create_time &&
				v.User == tmp.User && v.Content == tmp.Content{
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

func (self *CCHistory) Run() {

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
