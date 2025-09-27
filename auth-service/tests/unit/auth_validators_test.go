package unit

import (
	"auth-service/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type NameFieldTest struct {
	Name string `validate:"required,alpha,min=1"`
}

type SurnameFieldTest struct {
	Surname string `validate:"required,alpha,min=1"`
}

type GenderFieldTest struct {
	Gender string `validate:"required,oneof= male female"`
}

type PasswordFieldTest struct {
	Password string `validate:"required,min=8,alphanum,password"`
}

type EmailFieldTest struct {
	Email string `validate:"required,email,unique_email"`
}

type PhoneNumberFieldTest struct {
	PhoneNumber string `validate:"required,unique_phone,togolese_number"`
}

type FormatPhoneNumberTest struct {
	PhoneNumber string `validate:"required,togolese_number"`
}

func TestTogoleseNumberValidator(t *testing.T) {

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

func TestNameValidator(t *testing.T) {

	dto := NameFieldTest{Name: ""}
	dto2 := NameFieldTest{Name: "1234"}
	dto3 := NameFieldTest{Name: "john"}

	assert.Error(t, utils.ValidateRegistering.Struct(dto))
	assert.Error(t, utils.ValidateRegistering.Struct(dto2))
	assert.NoError(t, utils.ValidateRegistering.Struct(dto3))
}

func TestSurnameValidator(t *testing.T) {

	dto := SurnameFieldTest{Surname: ""}
	dto2 := SurnameFieldTest{Surname: "1234"}
	dto3 := SurnameFieldTest{Surname: "doe"}

	assert.Error(t, utils.ValidateRegistering.Struct(dto))
	assert.Error(t, utils.ValidateRegistering.Struct(dto2))
	assert.NoError(t, utils.ValidateRegistering.Struct(dto3))
}

func TestGenderValidator(t *testing.T) {

	dto := GenderFieldTest{Gender: ""}
	dto2 := GenderFieldTest{Gender: "male"}
	dto3 := GenderFieldTest{Gender: "female"}

	assert.Error(t, utils.ValidateRegistering.Struct(dto))
	assert.NoError(t, utils.ValidateRegistering.Struct(dto2))
	assert.NoError(t, utils.ValidateRegistering.Struct(dto3))
}

func TestPasswordValidator(t *testing.T) {

	dto := PasswordFieldTest{Password: ""}
	dto2 := PasswordFieldTest{Password: "abcd"}
	dto3 := PasswordFieldTest{Password: "abcdefghi"}
	dto4 := PasswordFieldTest{Password: "abdcdefg1234"}

	assert.Error(t, utils.ValidateRegistering.Struct(dto))
	assert.Error(t, utils.ValidateRegistering.Struct(dto2))
	assert.Error(t, utils.ValidateRegistering.Struct(dto3))
	assert.NoError(t, utils.ValidateRegistering.Struct(dto4))
}
