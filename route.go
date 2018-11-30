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

	users.POST("/", s.addUser)
	users.GET("/", s.getAllUser)
	users.GET("/:id", s.getUser)
	users.PUT("/:id", s.updateUser)

	return r
}
