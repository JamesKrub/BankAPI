package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

type UserInsert struct {
	ID        bson.ObjectId `json:"-"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

func (s *BankServiceImp) addUser(UserInsert) error {

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
	return u, nil
}

func (s *BankServiceImp) updateUserByID(id string) error {

	return nil
}

func (s *BankServiceImp) deleteUserByID(id string) error {

	return nil
}

func (s *BankServiceImp) addUserBankAccByID(string) error {

	return nil
}

func (s *BankServiceImp) getUserBankAccByID(string) error {

	return nil
}
