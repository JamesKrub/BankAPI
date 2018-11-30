package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

func (s *BankServiceImp) getAllListUser() ([]User, error) {
	var us []User

	err := s.db.C("users").Find(nil).All(&us)
	if err != nil {
		log.Fatal(err)
	}

	return us, nil
}
