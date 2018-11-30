package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoute(s *Server) *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	// bank := r.Group("/bankAccounts")
	admin := r.Group("/admin")

	admin.Use(gin.BasicAuth(gin.Accounts{
		"admin": "p",
	}))

	users.GET("/", s.getAllUser)
	users.POST("/", s.addUser)

	return r
}
