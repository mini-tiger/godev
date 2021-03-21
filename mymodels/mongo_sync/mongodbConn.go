package main

import (
	"godev/mymodels/mongo_sync/bussiness"
	"gopkg.in/mgo.v2"
	"time"
)

// xxx 不维护 其它方法见 GoDevEach\web框架\GinDemo\main.go

//type tableInfo struct {
//	databaseSession *mgo.Session
//}

type syncInterface interface {
	Init(sessionSrc, sessionDst *mgo.Session)
	DataAll() bool
	DiffSlice() bool
	Run()
	CollectionName() string
}

func createSession(ip, db string) *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{ip},
		Direct:    false,
		Timeout:   time.Second * 10,
		Database:  "cmdb",
		Source:    "admin",
		Username:  "root",
		Password:  "mongopass",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	return session
}

func main() {
	var ccapp bussiness.CCApp
	var cchistory bussiness.CCHistory
	var cchostbase bussiness.CCHostBase
	for {
		sessionSrc := createSession("192.168.130.17:27017", "cmdb")
		sessionDst := createSession("192.168.130.17:27017", "cmdb1")

		sessionSrc.SetMode(mgo.Monotonic, true)
		sessionDst.SetMode(mgo.Monotonic, true)
		//var i syncInterface
		Log := bussiness.Returnlog() // log

		tmp := make([]syncInterface, 0)
		tmp = append(tmp, &ccapp)
		tmp = append(tmp, &cchistory)
		tmp = append(tmp, &cchostbase)

		for _, t := range tmp {
			t.Init(sessionSrc, sessionDst)
			if !t.DataAll() {
				//log.Printf("collection:%s data empty","cc_ApplicationBase")
				Log.Info("collection:%s data empty", t.CollectionName())
				continue
			}
			if !t.DiffSlice() {
				Log.Info("collection:%s src dst Data same", t.CollectionName())
				continue
			}
			t.Run()

		}
		sessionSrc.Close()
		sessionDst.Close()
		time.Sleep(10 * time.Second)
	}

	// Optional. Switch the session to a monotonic behavior.

	//cc,err:=session.DB("test").CollectionNames()
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//fmt.Println(cc)

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
