package unit_test

import (
	"auth-service/internal/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "azerty1234$"
	hash, err := utils.HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword return an error %v", err)
	}

	if len(hash) == 0 {
		t.Fatalf("HashPassword must not be empty")
	}

	assert.NotEqual(t, password, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "azerty1234$"
	hash, err := utils.HashPassword(password)

	fmt.Print(hash)

	if err != nil {
		t.Fatalf("HashPassword return an error %v", err)
	}

	assert.True(t, utils.CheckPasswordHash(password, hash))
	assert.False(t, utils.CheckPasswordHash("wrong_password", hash))

}
