package main

type Server struct {
	bankService   BankService
	secretService SecretService
}

type SecretService interface {
	Insert(s *Secret) error
}

type BankService interface {
}

func main() {
	r := setupRoute(s)
}
