package integration

import (
	u "auth-service/internal/models"
	"auth-service/internal/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EmailFieldTest struct {
	Email string `validate:"required,email,unique_email"`
}

type PhoneNumberFieldTest struct {
	PhoneNumber string `validate:"required,unique_phone,togolese_number"`
}

type FormatPhoneNumberTest struct {
	PhoneNumber string `validate:"required,togolese_number"`
}

func TestUniqueEmailValidator(t *testing.T) {
	client := SetupDB(t)
	defer CloseDB(t, client)
	collection := client.Database("auth_test_db").Collection("users")

	utils.RegisterCustomValidators(collection)

	_, _ = collection.InsertOne(context.Background(), u.User{Email: "exists@example.com"})

	dto := EmailFieldTest{Email: "exists@example.com"}
	err := utils.ValidateRegistering.Struct(dto)
	assert.Error(t, err)

	dto2 := EmailFieldTest{Email: "unique@example.com"}
	err = utils.ValidateRegistering.Struct(dto2)
	assert.NoError(t, err)
}

func TestUniquePhoneNumberValidator(t *testing.T) {
	client := SetupDB(t)
	defer CloseDB(t, client)

	collection := client.Database("auth_test_db").Collection("users")

	utils.RegisterCustomValidators(collection)

	duplicatePhoneNumber := "90123456"
	uniquePhoneNumber := "71234567"

	_, _ = collection.InsertOne(context.Background(), u.User{PhoneNumber: duplicatePhoneNumber})

	dto := PhoneNumberFieldTest{PhoneNumber: duplicatePhoneNumber}
	err := utils.ValidateRegistering.Struct(dto)
	assert.Error(t, err)

	dto2 := PhoneNumberFieldTest{PhoneNumber: uniquePhoneNumber}
	err = utils.ValidateRegistering.Struct(dto2)
	assert.NoError(t, err)

}

func TestTogoleseNumberValidator(t *testing.T) {

	client := SetupDB(t)
	defer CloseDB(t, client)

	collection := client.Database("auth_test_db").Collection("users")

	utils.RegisterCustomValidators(collection)

	dto := FormatPhoneNumberTest{PhoneNumber: "94012345"}
	dto2 := FormatPhoneNumberTest{PhoneNumber: "990234"}
	dto3 := FormatPhoneNumberTest{PhoneNumber: "70123456"}

	err := utils.ValidateRegistering.Struct(dto)
	assert.Error(t, err)
	err = utils.ValidateRegistering.Struct(dto2)
	assert.Error(t, err)
	err = utils.ValidateRegistering.Struct(dto3)
	assert.NoError(t, err)
}
