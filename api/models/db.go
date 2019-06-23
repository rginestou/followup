package models

import (
	"gopkg.in/mgo.v2"
)

var userCol *mgo.Collection
var contactCol *mgo.Collection
var groupCol *mgo.Collection

func init() {
	// Open MongoDB server connection
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	userCol = session.DB("followup").C("user")
	contactCol = session.DB("followup").C("contact")
	groupCol = session.DB("followup").C("group")
}
