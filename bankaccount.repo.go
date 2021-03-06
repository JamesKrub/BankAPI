package main

import (
	"errors"
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

type Transfer struct {
	From   string `json:"from" bson:"balance"`
	To     string `json:"to" bson:"balance"`
	Amount int    `json:"amount"`
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

func (b *BankServiceImp) getBacnkAccDetailByBankAccID(id string) (UserBankAccount, error) {
	var acc UserBankAccount
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := b.db.C("accounts").Find(selector).One(&acc)
	if err != nil {
		fmt.Println("can't get bank acconut detail by bank account id")
		return acc, err
	}

	return acc, nil
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

	if t.Amount > acc.Balance {
		fmt.Println("The amount greater than balance")
		return errors.New("The amount greater than balance")
	}

	t.Amount = acc.Balance - t.Amount

	err = b.db.C("accounts").Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"balance": t.Amount}})
	if err != nil {
		fmt.Println("can't deposit")
		return err
	}
	return nil
}

func (b *BankServiceImp) transfer(t Transfer) error {
	var wd WithdrawDeposit
	wd.ID = bson.ObjectIdHex(t.From)
	wd.Amount = t.Amount
	err := b.withdrawByAccID(wd)
	if err != nil {
		fmt.Println("fail to minus From acc")
		return err
	}

	wd.ID = bson.ObjectIdHex(t.To)
	wd.Amount = t.Amount
	err = b.depositByAccID(wd)
	if err != nil {
		fmt.Println("fail to add To acc")
		return err
	}

	return nil
}

func (b *BankServiceImp) rollBack(id string, am int) error {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := b.db.C("accounts").Update(selector, bson.M{"$set": bson.M{"balance": am}})
	if err != nil {
		fmt.Println("fail to roll back")
		return err
	}
	return nil
}
