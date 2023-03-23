package database

import (
	"knowledge/api/model"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestInsertUser(t *testing.T) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := ConnectDatabase(connectionString)

	InsertUser(db, "testuser")
	user, err := GetUser(db, "testuser")

	if err != nil {
		t.Error(err)
	}

	if user == nil {
		t.Error("User not created")
	}
}

func TestUpdateUser(t *testing.T) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := ConnectDatabase(connectionString)

	user, _ := GetUser(db, "testuser")
	updatedUser := &model.User{
		Base: model.Base{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Username: "updated",
	}

	user, err := UpdateUser(db, updatedUser)

	if err != nil {
		t.Error(err)
	}

	if user == nil {
		t.Error("User not created")
	}
}

func TestDeleteUser(t *testing.T) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := ConnectDatabase(connectionString)

	err := DeleteUser(db, "updated")

	if err != nil {
		t.Error(err)
	}

	_, err = GetUser(db, "updated")

	if err == nil {
		t.Error("User not deleted")
	}
}
