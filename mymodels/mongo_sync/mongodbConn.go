package main

import (
	"fmt"
	"godev/mymodels/mongo_sync/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

type Person models.Person
type Person1 models.Person1

var M = map[string]interface{}{"people": Person{}}

func diffSlice(s1, s2 []interface{}) ([]Person, []Person, []Person) { //s1 源数据 s2 目标端数据

	if reflect.DeepEqual(s1, s2) {
		return []Person{}, []Person{}, []Person{}
	}
	tmpMap1 := make(map[bson.ObjectId]*Person, 0) // map[id]person{name,phone}   person不包含id, id当作KEY
	tmpMap2 := make(map[bson.ObjectId]*Person, 0)

	addSlice := make([]Person, 0)
	delSlice := make([]Person, 0)
	updateSlice := make([]Person, 0)

	for _, v := range s2 {
		p := &Person{}
		p.Name = v.Name
		p.Phone = v.Phone
		tmpMap2[v.Id_] = p
	}

	for _, v := range s1 {
		p := &Person{}
		p.Name = v.Name
		p.Phone = v.Phone
		tmpMap1[v.Id_] = p
	}

	for _, v := range s1 {
		if _, ok := tmpMap2[v.Id_]; ok { //源数据id在目标数据中有，应该比对其它相关字段
			if v.Name == tmpMap2[v.Id_].Name && v.Phone == tmpMap2[v.Id_].Phone { //其它字段都符合则跳过
				continue
			} else {
				updateSlice = append(updateSlice, v)
			}
		} else {
			addSlice = append(addSlice, v) //源数据有,目标没有
		}
	}

	for _, v := range s2 {
		if _, ok := tmpMap1[v.Id_]; ok { //目标有，源数据中有,不考虑目标库 同步到 源库
			continue
		} else {
			delSlice = append(delSlice, v) //目标有,源没有，应该从目标 删除
		}
	}
	return addSlice, delSlice, updateSlice

}

func Match(src, dst *mgo.Collection, t interface{}) {

	switch t.(type) {
	case Person:
		users := make([]Person, 0)
		users1 := make([]Person, 0)
	}

	src.Find(nil).All(&users)
	//c.Find(nil).Select(bson.M{"_id":1}).All(&users)

	dst.Find(nil).All(&users1)
	//c1.Find(nil).Select(bson.M{"_id":1}).All(&users1)

	//fmt.Println(users[0].Id_)

	//if len(users1) == len(users){ //判断长度相等意义 不大
	//	fmt.Println("相等")
	//}
	////a:=[]int{1}
	////a1:=[]int{1}
	if len(users) == 0 {
		return //源数据无，直接返回
	}
	add, del, update := diffSlice(users, users1)
	fmt.Println(add)
	fmt.Println(del)
	fmt.Println(update)
	for _, v := range add {
		p := Person{v.Id_, v.Name, v.Phone}
		c1.Insert(&p)
		//change := mgo.Change{
		//	Update:    bson.M{"$set": bson.M{p}},
		//	ReturnNew: false,
		//	Remove:    false,
		//	Upsert:    true,
		//}
		//p:=new(Person)
		//objectId := bson.ObjectIdHex("5bbb4d89159dd36a165b2417")
		//_, err = 	c.Find(nil).Apply(change, nil)
	}
	for _, v := range del {
		p := Person{v.Id_, v.Name, v.Phone}
		c1.Remove(bson.M{"_id": p.Id_})

	}

	for _, v := range update {
		//p := Person{v.Id_, v.Name, v.Phone}
		//
		//err:=c.Update(bson.M{"_id":v.Id_},p)
		//if err !=nil {
		//	fmt.Println(err)
		//}
		change := mgo.Change{
			Update:    bson.M{"$set": bson.M{"_id": v.Id_, "name": v.Name, "phone": v.Phone}},
			ReturnNew: false,
			Remove:    false,
			Upsert:    true,
		}
		//p := new(Person)
		//objectId := bson.ObjectId(v.Id_)
		_, err = c1.FindId(v.Id_).Apply(change, nil)

	}
}

func main() {
	session, err := mgo.Dial("1.119.132.143:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//cc,err:=session.DB("test").CollectionNames()
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//fmt.Println(cc)
	for k, v := range M {
		c := session.DB("test").C(k)
		c1 := session.DB("test1").C(k)
		switch k {
		case "people":
			s1 := make([]Person, 0)
			d1 := make([]Person, 0)
			Match(c, c1, Person())
		}

	}
	//change := mgo.Change{
	//	Update:    bson.M{"$set": bson.M{"phone": "1111111"}},
	//	ReturnNew: false,
	//	Remove:    false,
	//	Upsert:    true,
	//}
	////p:=new(Person)
	//objectId := bson.ObjectIdHex("5bbb4d89159dd36a165b2417")
	//_, err = 	c.FindId(objectId).Select(bson.M{"phone": 1}).Apply(change, nil)
	//
	//users1=make([]Person,0)
	//c.Find(nil).All(&users1)
	//fmt.Println(users1)
	//user := new(User)
	//p:=new(Person)
	//objectId := bson.ObjectIdHex("5bbb4d89159dd36a165b2417")
	//c.FindId(objectId).One(&p) // 直接通过ID查找
	//fmt.Println(p)
	//err = c.Insert(&Person{"superWang", "13478808311"},
	//	&Person{"David", "15040268074"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "superWang"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Name:", result.Name)
	//fmt.Println("Phone:", result.Phone)
}
