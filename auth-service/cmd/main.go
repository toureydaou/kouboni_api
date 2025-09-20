package main

import (
	"auth-service/internal/repository"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Database connection
	mongoUri := os.Getenv("MONGO_URL")
	if mongoUri == "" {
		log.Fatalf("Database URI not set")
	}
	repository.ConnectMongo(mongoUri)

	defer func() {
		err := repository.Client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Error while disconnecting from auth-service database %v", err)
		}
	}()

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("The auth service failed to start: %v \n", err)

	}

}
