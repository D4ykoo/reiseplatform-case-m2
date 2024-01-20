package test

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm/entities"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestFromEntityUser(t *testing.T) {
	user := entities.UserEntity{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Username:  "Test",
		Firstname: "Test",
		Lastname:  "Test",
		Email:     "test@test.test",
		Password:  "test",
		Salt:      "t3st",
	}

	modelUser := user.ToUser()

	if modelUser.Id != uint(1) {
		t.Fatalf(`ID: want %d - got: %d`, user.Model.ID, modelUser.Id)
	}

	if modelUser.Username != "Test" {
		t.Fatalf(`Username: want %s - got: %s`, user.Username, modelUser.Username)
	}

	if modelUser.Firstname != "Test" {
		t.Fatalf(`Firstname: want %s - got: %s`, user.Firstname, modelUser.Firstname)
	}

	if modelUser.Lastname != "Test" {
		t.Fatalf(`Lastname: want %s - got: %s`, user.Lastname, modelUser.Lastname)
	}

	if modelUser.Email != "test@test.test" {
		t.Fatalf(`Email: want %s - got: %s`, user.Email, modelUser.Email)
	}

	if modelUser.Password != "test" {
		t.Fatalf(`Password: want %s - got: %s`, user.Password, modelUser.Password)
	}
}
