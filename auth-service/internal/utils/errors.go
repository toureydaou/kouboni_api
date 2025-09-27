package utils

import "errors"

var (
	ErrorEmailExists       = errors.New("email already in use")
	ErrorPhoneNumberExists = errors.New("phone number already in use")
)
