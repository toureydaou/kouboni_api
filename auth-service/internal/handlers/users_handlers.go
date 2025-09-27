package handlers

import (
	u "auth-service/internal/models"
	"auth-service/internal/services"
	"context"
	"log"

	"auth-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService *services.AuthService
}

func NewUserHandler(authService *services.AuthService) *UserHandler {
	return &UserHandler{authService: authService}
}

func (h *UserHandler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Working !")
}

func (h *UserHandler) RegisterHandler(c *gin.Context) {

	var input u.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Binding error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateRegistering.Struct(input); err != nil {
		log.Printf("Validation error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.authService.RegisterUser(context.Background(), input)

	if err != nil {

		switch err {
		case utils.ErrorEmailExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case utils.ErrorPhoneNumberExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			log.Printf("Register internal error %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"response": userResponse})
}
