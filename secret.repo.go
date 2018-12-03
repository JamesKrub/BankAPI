package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Secret struct {
	ID  bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Key string        `json:"key" bson:"key"`
}

func (b *BankServiceImp) addSecret(s Secret) error {
	err := b.db.C("secrets").Insert(s)
	if err != nil {
		fmt.Println("can't insert secret")
		return err
	}
	return nil
}

func (b *BankServiceImp) getSecret(s string) error {
	selector := bson.M{"key": s}
	_, err := b.db.C("secrets").Find(selector).Count()
	if err != nil {
		fmt.Println("can't count selected secret")
		return err
	}
	return nil
}
