package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAllUser(c *gin.Context) {
	us, err := s.bankService.getAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[getAllUser] get user to db got error: %s", err),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
		"data":   us,
	})
}

func (s *Server) addUser(c *gin.Context) {
	var u UserInsert
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json has wrong params: %s", err),
		})
		return
	}

	err = s.bankService.addUser(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("insert User got error: %s", err),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"object": "success",
	})
}
