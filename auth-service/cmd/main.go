package main

import (
	"auth-service/internal/repository"
	"context"
	"log"
	"os"
	"time"

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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := repository.Client.Disconnect(ctx); err != nil {
			log.Fatalf("Error while disconnecting from auth-service database %v", err)
		}
	}()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("The auth service failed to start: %v \n", err)

	}

}
