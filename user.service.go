package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) getAllUser(c *gin.Context) {
	s.bankService.getAllListUser()
}
