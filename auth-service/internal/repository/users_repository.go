package repository

import (
	u "auth-service/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user u.User) error
	FindUserByEmail(ctx context.Context, email string) (user *u.User, error error)
	FindUserByPhoneNumber(ctx context.Context, email string) (user *u.User, error error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user u.User) error {

	_, err := r.collection.InsertOne(ctx, user)
	return err

}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*u.User, error) {

	var user *u.User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (*u.User, error) {
	var user *u.User

	err := r.collection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
