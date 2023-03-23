package database

import (
	"knowledge/api/model"
	"log"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}

	err := db.Create(user).Error
	if err != nil {
		log.Panic("Unable to create user.")
	}

	return user, err
}

func GetUser(db *gorm.DB, username string) (*model.User, error) {
	var user model.User

	err := db.Where("username =?", username).First(&user).Error

	return &user, err
}

func GetUsers(db *gorm.DB) ([]*model.User, error) {
	var users []*model.User

	err := db.Find(&users).Error

	return users, err
}

func UpdateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Save(user).Error

	return user, err
}

func DeleteUser(db *gorm.DB, username string) error {
	var user model.User

	err := db.Where("username =?", username).Delete(&user).Error

	return err
}
