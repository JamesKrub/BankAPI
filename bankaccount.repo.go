package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type UserBankAccount struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AccountNumber string        `json:"account_number" bson:"account_number"`
	Name          string        `json:"name" bson:"name"`
	UserID        string        `json:"user_id" bson:"user_id"`
	Balance       int           `json:"balance" bson:"balance"`
}

type UserBankAccountInsert struct {
	AccountNumber string `json:"account_number" bson:"account_number"`
	Name          string `json:"name" bson:"name"`
	UserID        string `json:"user_id" bson:"user_id"`
	Balance       int    `json:"balance" bson:"balance"`
}

func (s *BankServiceImp) addBankAccByUserID(ac UserBankAccountInsert) error {
	err := s.db.C("accounts").Insert(ac)
	if err != nil {
		fmt.Println("can't insert user bank acconut")
		return err
	}
	return nil
}

func (s *BankServiceImp) countBankAccByBankAccID(id string) (int, error) {
	selector := bson.M{"account_number": id}
	count, err := s.db.C("users").Find(selector).Count()
	if err != nil {
		fmt.Println("can't insert user bank acconut")
		return 0, err
	}
	return count, nil
}

func (s *BankServiceImp) getBankAccByUserID(string) error {

	return nil
}
