package api

import (
	"auth-service/internal/handlers"
	"auth-service/internal/repository"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userRepo repository.UserRepository) *gin.Engine {

	userService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()

	authRoutes := router.Group("/auth")

	authRoutes.POST("/register", userHandler.RegisterHandler)
	authRoutes.GET("/health", userHandler.HealthHandler)

	return router
}
