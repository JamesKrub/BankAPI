package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

func (s *BankServiceImp) getAllListUser() ([]User, error) {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/bank")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var us []User
	c := session.DB("bank").C("users")
	err = c.Find(nil).All(&us)
	if err != nil {
		log.Fatal(err)
	}

	return us, nil
}
