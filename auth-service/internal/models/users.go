package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Surname     string             `bson:"surname,omitempty"`
	Gender      string             `bson:"gender,omitempty"`
	Email       string             `bson:"email,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	Password    string             `bson:"password,omitempty"`
}

type UserResponse struct {
	ID          string `json:"id"`
	FullName    string `json:"fullname"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserRegister struct {
	Name        string `json:"name" validate:"required,alpha,min=1"`
	Surname     string `json:"surname" validate:"required,alpha,min=1"`
	Email       string `json:"email" validate:"email"`
	Gender      string `json:"gender" validate:"required,oneof= male female"`
	PhoneNumber string `json:"telephone" validate:"required,togolese_number"`
	Password    string `json:"password" validate:"required,min=8,password"`
}
