package e2e

import (
	"auth-service/api"
	u "auth-service/internal/models"
	"auth-service/internal/repository"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USER_NOT_FOUND = "User not found"
const USER_COLLECTION = "users"
const ERROR_USER_QUERYING = "Error while quering users"
const TEST_DB = "auth_test_db"

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

var userRegister *u.UserRegister
var userRegister2 *u.UserRegister

func initUser() {
	userRegister = &u.UserRegister{
		Name:        "John",
		Surname:     "Doe",
		Gender:      "male",
		Email:       "johndoe@email.com",
		PhoneNumber: "90123456",
		Password:    "azerty123",
	}

	userRegister2 = &u.UserRegister{
		Name:        "Suzanne",
		Surname:     "Doe",
		Gender:      "female",
		Email:       "suzanne@email.com",
		PhoneNumber: "90123456",
		Password:    "azerty123",
	}

}

const REGISTER_ENDPOINT = "/auth/register"
const HEADER_CONTENT_TYPE = "application/json"

func TestRegisterSuccess(t *testing.T) {
	db := SetupDB(t)
	defer CloseDB(t, db)

	repo := repository.NewUserRepository(db.Database(TEST_DB), USER_COLLECTION)
	r := api.SetupRoutes(repo)
	ts := httptest.NewServer(r)

	initUser()

	body, err := json.Marshal(*userRegister)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL+REGISTER_ENDPOINT, HEADER_CONTENT_TYPE, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Error(err)
		}
	}()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

}

func TestRegisterDuplicatedEmail(t *testing.T) {
	db := SetupDB(t)
	defer CloseDB(t, db)

	repo := repository.NewUserRepository(db.Database(TEST_DB), USER_COLLECTION)

	r := api.SetupRoutes(repo)

	ts := httptest.NewServer(r)

	initUser()

	body, err := json.Marshal(*userRegister)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL+REGISTER_ENDPOINT, HEADER_CONTENT_TYPE, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Error(err)
		}
	}()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL+REGISTER_ENDPOINT, HEADER_CONTENT_TYPE, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusConflict, resp.StatusCode)

}

func TestRegisterDuplicatedPhoneNumber(t *testing.T) {
	db := SetupDB(t)
	defer CloseDB(t, db)

	repo := repository.NewUserRepository(db.Database(TEST_DB), USER_COLLECTION)

	r := api.SetupRoutes(repo)

	ts := httptest.NewServer(r)

	initUser()

	body, err := json.Marshal(*userRegister)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL+REGISTER_ENDPOINT, HEADER_CONTENT_TYPE, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Error(err)
		}
	}()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	body, err = json.Marshal(*userRegister)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.Post(ts.URL+REGISTER_ENDPOINT, HEADER_CONTENT_TYPE, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusConflict, resp.StatusCode)

}
