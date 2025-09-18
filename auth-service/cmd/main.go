package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "auth service",
		})
	})
	router.Run(":8080")

}
