package main

import (
	"os"
)

type Server struct {
	bankService   BankService
	secretService SecretService
}

type SecretService interface {
}

type BankService interface {
	getAllListUser() ([]User, error)
}

func main() {
	s := &Server{
		bankService:   &BankServiceImp{},
		secretService: &SecertServiceImp{},
	}
	r := setupRoute(s)
	r.Run(":" + os.Getenv("PORT"))
}

type BankServiceImp struct {
}

type SecertServiceImp struct {
}
