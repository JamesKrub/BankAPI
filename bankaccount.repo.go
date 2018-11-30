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

type WithdrawDeposit struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount int           `json:"amount" bson:"balance"`
}

type Ttransfer struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	From int           `json:"from" bson:"balance"`
	To   int           `json:"to" bson:"balance"`
}

func (b *BankServiceImp) addBankAccByUserID(ac UserBankAccountInsert) error {
	err := b.db.C("accounts").Insert(ac)
	if err != nil {
		fmt.Println("can't insert user bank acconut")
		return err
	}
	return nil
}

func (b *BankServiceImp) countBankAccByBankAccID(id string) (int, error) {
	selector := bson.M{"account_number": id}
	count, err := b.db.C("accounts").Find(selector).Count()
	if err != nil {
		fmt.Println("can't insert user bank acconut")
		return 0, err
	}
	return count, nil
}

func (b *BankServiceImp) getBankAccByUserID(id string) ([]UserBankAccount, error) {
	var accs []UserBankAccount
	selector := bson.M{"user_id": id}
	err := b.db.C("accounts").Find(selector).All(&accs)
	if err != nil {
		fmt.Println("can't get bank acconut by user id")
		return accs, err
	}
	return accs, nil
}

func (b *BankServiceImp) deleteBankAccByBankAccID(id string) error {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := b.db.C("accounts").Remove(selector)
	if err != nil {
		fmt.Println("can't delelte bank acconut by id")
		return err
	}
	return nil
}

func (b *BankServiceImp) depositByAccID(t WithdrawDeposit) error {
	var acc UserBankAccount
	err := b.db.C("accounts").Find(bson.M{"_id": t.ID}).One(&acc)
	if err != nil {
		fmt.Println("can't get acc detail")
		return err
	}

	t.Amount = t.Amount + acc.Balance

	err = b.db.C("accounts").Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"balance": t.Amount}})
	if err != nil {
		fmt.Println("can't deposit")
		return err
	}
	return nil
}

func (b *BankServiceImp) withdrawByAccID(t WithdrawDeposit) error {
	var acc UserBankAccount
	err := b.db.C("accounts").Find(bson.M{"_id": t.ID}).One(&acc)
	if err != nil {
		fmt.Println("can't get acc detail")
		return err
	}

	t.Amount = acc.Balance - t.Amount

	err = b.db.C("accounts").Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"balance": t.Amount}})
	if err != nil {
		fmt.Println("can't deposit")
		return err
	}
	return nil
}

func (b *BankServiceImp) transfer(t WithdrawDeposit) error {
	return nil
}
