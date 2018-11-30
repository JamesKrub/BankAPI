package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

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

func (s *Server) getUser(c *gin.Context) {
	id := c.Param("id")
	u, err := s.bankService.getUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[getUser] getUserByID got error: %v", err),
		})
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
		"data":   u,
	})
}

func (s *Server) updateUser(c *gin.Context) {
	var u UserUpdate
	id := c.Param("id")

	c.ShouldBindJSON(&u)
	u.ID = bson.ObjectIdHex(id)
	err := s.bankService.updateUserByID(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("[updateUser] updateUserByID got error: %v", err),
		})
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"object": "success",
	})
}
