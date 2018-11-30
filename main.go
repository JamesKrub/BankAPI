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

type SecretService interface {
}

type BankService interface {
	getAllListUser() ([]User, error)
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

type BankServiceImp struct {
	db *mgo.Database
}

type SecertServiceImp struct {
	db *mgo.Database
}
