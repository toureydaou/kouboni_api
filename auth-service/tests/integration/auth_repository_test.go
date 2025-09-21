package integration

import (
	"context"
	"testing"
	"time"

	u "auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(t *testing.T) (*mongo.Collection, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/auth_test_db"))

	if err != nil {
		t.Fatalf("Failed to connect to database %v", err)
	}

	database := client.Database("auth_test_db")

	collection := database.Collection("users")

	if err := collection.Drop(ctx); err != nil {
		t.Fatalf("Failed to drop collection")
	}

	return collection, client
}

func CloseDB(t *testing.T, client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		t.Fatalf("Error when disconnecting from database %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	// test the insertion of an user

	usersCollection, client := SetupDB(t)

	defer CloseDB(t, client)

	user := u.User{
		FirstName:   "John",
		Surname:     "Doe",
		Email:       "johndoe@email.com",
		PhoneNumber: "90123456",
		Password:    utils.HashPassword("azerty$123"),
	}

	if err := repository.CreateUser(user); err != nil {
		t.Fatalf("Error while inserting user %v", err)
	}

	var foundUser u.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := usersCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
		t.Fatalf("Error while quering users %v", err)
	}

	assert.Equal(t, foundUser.Email, user.Email)

}
