package main

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

func getAllListUser() ([]User, error) {
	var us []User

	return us, nil
}
