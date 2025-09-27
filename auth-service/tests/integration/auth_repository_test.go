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

const USER_NOT_FOUND = "User not found"
const USER_COLLECTION = "users"
const ERROR_USER_QUERYING = "Error while quering users"
const TEST_DB = "auth_test_db"

var user *u.User

func initUser() {
	user = &u.User{
		Name:        "John",
		Surname:     "Doe",
		Gender:      "M",
		Email:       "johndoe@email.com",
		PhoneNumber: "90123456",
		Password:    utils.HashPassword("azerty$123"),
	}

}

func SetupDB(t *testing.T) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/auth_test_db"))
	if err != nil {
		t.Fatalf("Failed to connect to database %v", err)
	}

	database := client.Database(TEST_DB)
	collection := database.Collection(USER_COLLECTION)
	if err := collection.Drop(ctx); err != nil {
		t.Fatalf("Failed to drop collection %v", err)
	}

	return client
}

func CloseDB(t *testing.T, client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		t.Fatalf("Error when disconnecting from database %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	client := SetupDB(t)
	defer CloseDB(t, client)

	initUser()

	repo := repository.NewUserRepository(client.Database(TEST_DB), USER_COLLECTION)

	if err := repo.CreateUser(context.Background(), *user); err != nil {
		t.Fatalf("Error while inserting user %v", err)
	}

	var foundUser u.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Database(TEST_DB).Collection(USER_COLLECTION).FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
		if err == mongo.ErrNoDocuments {
			t.Error(USER_NOT_FOUND)
		}
		t.Fatalf(ERROR_USER_QUERYING+" %v", err)
	}
	assert.Equal(t, foundUser.Email, user.Email)
}

func TestFindUserByEmail(t *testing.T) {
	client := SetupDB(t)
	defer CloseDB(t, client)

	initUser()

	db := client.Database(TEST_DB)

	repo := repository.NewUserRepository(db, USER_COLLECTION)

	if err := repo.CreateUser(context.Background(), *user); err != nil {
		if err == mongo.ErrNoDocuments {
			t.Error(USER_NOT_FOUND)
		}
		t.Fatalf(ERROR_USER_QUERYING+" %v", err)
	}

	userFound, err := repo.FindUserByEmail(context.Background(), user.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			t.Error(USER_NOT_FOUND)
		}
		t.Fatalf(ERROR_USER_QUERYING+" %v", err)
	}

	assert.Equal(t, userFound.Email, user.Email)
}

func TestFindUserByPhoneNumber(t *testing.T) {
	client := SetupDB(t)
	defer CloseDB(t, client)

	initUser()
	db := client.Database(TEST_DB)

	repo := repository.NewUserRepository(db, USER_COLLECTION)

	if err := repo.CreateUser(context.Background(), *user); err != nil {
		if err == mongo.ErrNoDocuments {
			t.Error(USER_NOT_FOUND)
		}
		t.Fatalf(ERROR_USER_QUERYING+" %v", err)
	}

	userFound, err := repo.FindUserByPhoneNumber(context.Background(), user.PhoneNumber)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			t.Error(USER_NOT_FOUND)
		}
		t.Fatalf(ERROR_USER_QUERYING+" %v", err)
	}

	assert.Equal(t, userFound.PhoneNumber, user.PhoneNumber)
}
