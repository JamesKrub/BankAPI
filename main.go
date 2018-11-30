package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	admin := r.Group("/admin/")
}
