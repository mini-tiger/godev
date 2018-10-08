package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"reflect"
	"time"
)

type Person struct {
	Id_   bson.ObjectId `bson:"_id"`
	Name  string
	Phone string
}

func diffSlice(s1, s2 []Person) ([]Person, []Person) { //s1 源数据 s2 目标端数据
	if reflect.DeepEqual(s1, s2) {
		return []Person{}, []Person{}
	}
	tmpMap1 := make(map[Person]struct{}, 0)
	tmpMap2 := make(map[Person]struct{}, 0)

	addSlice := make([]Person, 0)
	delSlice := make([]Person, 0)
	for _, v := range s2 {
		tmpMap2[v] = struct{}{}
	}

	for _, v := range s1 {
		tmpMap1[v] = struct{}{}
	}

	for _, v := range s1 {
		if _, ok := tmpMap2[v]; ok { //源数据id在目标数据中有 应该比对其它字段
			continue
		} else {
			addSlice = append(addSlice, v) //源数据有,目标没有，应该添加到 目标
		}
	}

	for _, v := range s2 {
		if _, ok := tmpMap1[v]; ok { //目标id有，源数据中有 应该比对其它字段。上面或者本循环检验
			continue
		} else {
			delSlice = append(delSlice, v) //目录有,源没有，应该从目标 删除
		}
	}
	return addSlice, delSlice

}

func main() {
	session, err := mgo.Dial("1.119.132.143:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	p := Person{bson.ObjectIdHex(fmt.Sprintf("5bbb4d89159dd36a165b240%d", 1)), "superWang", "13478808311"}
	for k, v := range p {
		fmt.Println(k, v)
	}
	//cc,err:=session.DB("test").CollectionNames()
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//fmt.Println(cc)
	c := session.DB("test").C("people")
	for i := 0; i < 10; i++ {
		err = c.Insert(&Person{bson.ObjectIdHex(fmt.Sprintf("5bbb4d89159dd36a165b240%d", i)), "superWang", "13478808311"})
		time.Sleep(2 * time.Second)
	}

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
