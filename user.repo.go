package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

type UserInsert struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type UserUpdate struct {
	ID        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

func (s *BankServiceImp) addUser(u UserInsert) error {
	err := s.db.C("users").Insert(u)
	if err != nil {
		fmt.Println("can't Insert to db")
		return err
	}
	return nil
}

func (s *BankServiceImp) getAllUser() ([]User, error) {
	var us []User
	err := s.db.C("users").Find(nil).All(&us)
	if err != nil {
		log.Fatal(err)
	}

	return us, nil
}

func (s *BankServiceImp) getUserByID(id string) (User, error) {
	var u User
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := s.db.C("users").Find(selector).One(&u)
	if err != nil {
		fmt.Println("can't get user data")
		return u, err
	}

	return u, nil
}

func (s *BankServiceImp) updateUserByID(u UserUpdate) error {
	err := s.db.C("users").Update(bson.M{"_id": u.ID}, bson.M{"$set": u})
	if err != nil {
		return err
	}
	return nil
}

func (s *BankServiceImp) deleteUserByID(id string) error {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := s.db.C("users").Remove(selector)
	if err != nil {
		fmt.Println("can't delete user")
		return err
	}
	return nil
}

func (s *BankServiceImp) countUserByID(id string) (int, error) {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	count, err := s.db.C("users").Find(selector).Count()
	if err != nil {
		fmt.Println("can't count user")
		return 0, err
	}
	return count, nil
}
