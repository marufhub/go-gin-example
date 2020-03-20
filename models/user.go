package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// User model
	User struct {
		gorm.Model
		UserID   int    `json:"user_id"`
		UserName string `json:"user_name"`
	}
)

// GetAllUsers return the list of all users
func GetAllUsers() []*User {
	list := []*User{}
	GetDB().Find(&list)

	return list
}

// GetUserByID return a user by id
func GetUserByID(id int64) *User {
	user := &User{}
	if GetDB().Where("user_id = ?", id).First(user); user.UserID > 0 {
		return user
	}
	return nil
}
