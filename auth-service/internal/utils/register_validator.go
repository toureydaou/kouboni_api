package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
)

const ERROR_VALIDATOR_REGISTERING = "Error while registering validator "
const TOGOLESE_NUMBER_VALIDATOR = "togolese_number"
const PASSWORD_VALIDATOR = "password"

var ValidateRegistering = validator.New()

func init() {
	if err := ValidateRegistering.RegisterValidation(TOGOLESE_NUMBER_VALIDATOR, TogoleseNumberValidator); err != nil {
		log.Printf(ERROR_VALIDATOR_REGISTERING+TOGOLESE_NUMBER_VALIDATOR+" %v", err)
	}
	if err := ValidateRegistering.RegisterValidation(PASSWORD_VALIDATOR, PasswordValidator); err != nil {
		log.Printf(ERROR_VALIDATOR_REGISTERING+PASSWORD_VALIDATOR+" %v", err)
	}
}
