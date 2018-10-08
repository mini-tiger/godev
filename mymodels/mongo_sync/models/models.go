package models

import "gopkg.in/mgo.v2/bson"

type Person struct {
	Id_   bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
}

type Person1 struct {
	Id_   bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
}
