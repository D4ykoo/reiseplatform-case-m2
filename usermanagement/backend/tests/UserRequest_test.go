package test

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api/dto"
	"testing"
)

func TestResponseToModelUser(t *testing.T) {
	createUserRequest := dto.CreateUserRequest{
		Username:  "Test",
		Firstname: "Test",
		Lastname:  "Test",
		Email:     "test@test.test",
		Password:  "test",
	}

	modelUser := createUserRequest.ToUser()

	if modelUser.Username != "Test" {
		t.Fatalf(`Username: want %s - got: %s`, createUserRequest.Username, modelUser.Username)
	}

	if modelUser.Firstname != "Test" {
		t.Fatalf(`Firstname: want %s - got: %s`, createUserRequest.Firstname, modelUser.Firstname)
	}

	if modelUser.Lastname != "Test" {
		t.Fatalf(`Lastname: want %s - got: %s`, createUserRequest.Lastname, modelUser.Lastname)
	}

	if modelUser.Email != "test@test.test" {
		t.Fatalf(`Email: want %s - got: %s`, createUserRequest.Email, modelUser.Email)
	}

	if modelUser.Password != "test" {
		t.Fatalf(`Password: want %s - got: %s`, createUserRequest.Password, modelUser.Password)
	}
}
