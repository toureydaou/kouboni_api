package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("The auth service failed to start: %v \n", err)
	}

}
