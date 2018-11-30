package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
}

type UserInsert struct {
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
	err := s.db.C("users").Find(nil).One(&u)
	if err != nil {
		fmt.Println("can't get user data")
		return u, err
	}

	return u, nil
}

func (s *BankServiceImp) updateUserByID(id string) error {

	return nil
}

func (s *BankServiceImp) deleteUserByID(id string) error {

	return nil
}

func (s *BankServiceImp) addBankAccByUserID(string) error {

	return nil
}

func (s *BankServiceImp) getBankAccByUserID(string) error {

	return nil
}
