package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"first_name"`
	Surname     string             `bson:"surname"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
	Password    string             `bson:"password"`
}

type UserResponse struct {
	ID          string `json:"id"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
