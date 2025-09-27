package unit_test

import (
	m "auth-service/internal/models"
	"auth-service/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHashPassword(t *testing.T) {
	password := "azerty1234$"
	hash := utils.HashPassword(password)

	if len(hash) == 0 {
		t.Fatalf("HashPassword must not be empty")
	}

	assert.NotEqual(t, password, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "azerty1234$"
	hash := utils.HashPassword(password)

	assert.True(t, utils.CheckPasswordHash(password, hash))
	assert.False(t, utils.CheckPasswordHash("wrong_password", hash))

}

func TestToUserResponse(t *testing.T) {
	userId := primitive.NewObjectID()

	user := m.User{
		ID:          userId,
		Name:        "John",
		Surname:     "Doe",
		Gender:      "male",
		Email:       "johndoe@email.com",
		PhoneNumber: "90123456",
		Password:    utils.HashPassword("$aze1ty*"),
	}

	expectedUserResponse := m.UserResponse{
		ID:          userId.Hex(),
		FullName:    "John Doe",
		Gender:      "male",
		Email:       "johndoe@email.com",
		PhoneNumber: "90123456",
	}

	userResponse := utils.ToUserResponse(user)

	assert.Equal(t, *userResponse, expectedUserResponse)

}
