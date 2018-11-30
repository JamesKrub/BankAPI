package main

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	db            *mgo.Database
	bankService   BankService
	secretService SecretService
}

type BankServiceImp struct {
	db *mgo.Database
}

type SecertServiceImp struct {
	db *mgo.Database
}

type SecretService interface {
}

type BankService interface {
	addUser(UserInsert) error
	getAllUser() ([]User, error)
	getUserByID(string) (User, error)
	updateUserByID(UserUpdate) error
	deleteUserByID(string) error
	countUserByID(string) (int, error)
	addBankAccByUserID(UserBankAccountInsert) error
	countBankAccByBankAccID(string) (int, error)
	getBankAccByUserID(string) ([]UserBankAccount, error)
	deleteBankAccByBankAccID(string) error
	depositByAccID(WithdrawDeposit) error
	withdrawByAccID(WithdrawDeposit) error
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/bank")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("bank")

	s := &Server{
		db: session.DB("bank"),
		bankService: &BankServiceImp{
			db: db,
		},
		secretService: &SecertServiceImp{
			db: db,
		},
	}
	r := setupRoute(s)
	r.Run(":" + os.Getenv("PORT"))
}
