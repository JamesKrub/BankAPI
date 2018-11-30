package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) addBankAccount(c *gin.Context) {
	var acc UserBankAccountInsert
	id := c.Param("id")
	acc.UserID = id

	count, err := s.bankService.countUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[addBankAccount] countUserByID got error: %v", err),
		})
		return
	}

	if count <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "Request ID doesn't exist",
		})
		return
	}

	err = s.bankService.addBankAccByUserID(acc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[addBankAccount] addBankAccByUserID got error: %v", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}
