package utils

import (
	"context"
	"log"
	"regexp"
	"time"

	u "auth-service/internal/models"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	re := regexp.MustCompile(`[0-9]`)

	return re.MatchString(password)
}

func TogoleseNumberValidator(fl validator.FieldLevel) bool {

	phoneNumber := fl.Field().String()

	re := regexp.MustCompile(`^((9|7)([0-3]|[6-9]))\d{6}$`)

	return re.MatchString(phoneNumber)
}

func UniqueEmailValidator(collection *mongo.Collection) validator.Func {
	return func(fl validator.FieldLevel) bool {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		email := fl.Field().String()
		var user u.User

		err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return true
			}

			log.Printf("Error querying user by email: %v", err)
			return false
		}

		return false
	}
}

func UniquePhoneNumberValidator(collection *mongo.Collection) validator.Func {
	return func(fl validator.FieldLevel) bool {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var user u.User

		phoneNumber := fl.Field().String()

		err := collection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return true
			}
			log.Printf("Error querying user by phone number: %v", err)
			return false
		}
		return false
	}
}
