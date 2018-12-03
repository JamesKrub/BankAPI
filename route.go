package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoute(s *Server) *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	bank := r.Group("/bankAccounts")
	transfer := r.Group("/transfer")
	admin := r.Group("/admin")

	admin.Use(gin.BasicAuth(gin.Accounts{
		"admin": "p",
	}))

	admin.POST("/secrets", s.addSecrets)

	users.Use(s.authBankAPI)
	users.PUT("/:id", s.updateUser)
	users.DELETE("/:id", s.deleteUser)
	users.POST("/:id/bankAccounts", s.addBankAccount)
	users.GET("/:id/bankAccounts", s.getBankAccount)
	bank.DELETE("/:id", s.deleteBankAcconut)
	bank.PUT("/:id/deposit", s.deposit)
	bank.PUT("/:id/withdraw", s.withdraw)
	transfer.POST("/", s.transfer)

	return r
}
