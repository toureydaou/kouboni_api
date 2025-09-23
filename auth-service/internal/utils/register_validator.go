package utils

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var ValidateRegistering = validator.New()

func RegisterCustomValidators(collection *mongo.Collection) {
	ValidateRegistering.RegisterValidation("unique_email", UniqueEmailValidator(collection))
	ValidateRegistering.RegisterValidation("unique_phone", UniquePhoneNumberValidator(collection))
	ValidateRegistering.RegisterValidation("togolese_number", TogoleseNumberValidator)
	ValidateRegistering.RegisterValidation("password", PasswordValidator)
}
