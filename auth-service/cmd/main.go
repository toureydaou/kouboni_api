package main

import (
	"auth-service/api"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"context"
	"log"
	"os"
	"time"
)

func main() {

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		log.Fatalf("Database URI not set")
	}
	client := repository.ConnectMongo(mongoUri)
	db := client.Database(utils.AUTH_DB)

	userRepo := repository.NewUserRepository(db, utils.USER_COLLECTION)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := repository.Client.Disconnect(ctx); err != nil {
			log.Fatalf("Error while disconnecting from auth-service database %v", err)
		}
	}()

	router := api.SetupRoutes(userRepo)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("The auth service failed to start: %v \n", err)

	}

}
