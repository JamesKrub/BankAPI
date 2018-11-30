package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) addBankAccount(c *gin.Context) {
	var acc UserBankAccountInsert
	id := c.Param("id")
	c.ShouldBindJSON(&acc)
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

	count, err = s.bankService.countBankAccByBankAccID(acc.AccountNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[addBankAccount] countBankAccByBankAccID got error: %v", err),
		})
		return
	}

	if count != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "Duplicate bank account number",
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

func (s *Server) getBankAccount(c *gin.Context) {
	id := c.Param("id")
	accs, err := s.bankService.getBankAccByUserID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[getBankAccount] getBankAccByUserID got error: %v", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
		"data":   accs,
	})
}

func (s *Server) deleteBankAcconut(c *gin.Context) {
	id := c.Param("id")
	err := s.bankService.deleteBankAccByBankAccID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[deleteBankAcconut] delete bank account got error: %v", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}

func (s *Server) deposit(c *gin.Context) {
	// id := c.Param("id")
}
