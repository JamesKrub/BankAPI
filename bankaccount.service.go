package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (s *Server) authBankAPI(c *gin.Context) {
	user, _, ok := c.Request.BasicAuth()
	if ok {
		err := s.bankService.getSecret(user)
		if err == nil {
			return
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (s *Server) addSecrets(c *gin.Context) {
	var secret Secret
	err := c.ShouldBindJSON(&secret)
	if err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"object":  "error",
				"message": fmt.Sprintf("[addSecrets] json parse got error: %v", err),
			})
			return
		}
	}

	err = s.bankService.addSecret(secret)
	if err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"object":  "error",
				"message": fmt.Sprintf("[addSecrets] addSecret got error: %v", err),
			})
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}

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
	id := c.Param("id")
	var t WithdrawDeposit
	c.ShouldBindJSON(&t)
	t.ID = bson.ObjectIdHex(id)

	err := s.bankService.depositByAccID(t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[deposit] deposit  got error: %v", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}

func (s *Server) withdraw(c *gin.Context) {
	id := c.Param("id")
	var t WithdrawDeposit
	c.ShouldBindJSON(&t)
	t.ID = bson.ObjectIdHex(id)

	err := s.bankService.withdrawByAccID(t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[deposit] deposit  got error: %v", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}

func (s *Server) transfer(c *gin.Context) {
	var t Transfer
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[transfer] json parse got error: %v", err),
		})
		return
	}
	r, err := s.bankService.getBacnkAccDetailByBankAccID(t.From)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[transfer] get bank acc detail by From ID got error: %v", err),
		})
		return
	}
	from := r.Balance

	r, err = s.bankService.getBacnkAccDetailByBankAccID(t.To)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[transfer] get bank acc detail by To ID got error: %v", err),
		})
		return
	}
	to := r.Balance

	err = s.bankService.transfer(t)
	if err != nil {
		err = s.bankService.rollBack(t.From, from)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"object":  "error",
				"message": fmt.Sprintf("[transfer] panic!! [From]recovery got error: %v", err),
			})
		}

		err = s.bankService.rollBack(t.To, to)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"object":  "error",
				"message": fmt.Sprintf("[transfer] panic!! [To]recovery got error: %v", err),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[transfer] transfer monety got error: %v {system roll back}", err),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}
