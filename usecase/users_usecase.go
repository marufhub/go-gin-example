package usecase

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/db"
	"github.com/marufhub/go-gin-example/model"
)

// GetUsers return the list of all users
func GetUsers() {
	var users []model.User
	db := db.GetDB()
	db.Find(&users)

	return users, nil

}

// GetUser returns a user for the user_id
func GetUser(user_id string) {
	var user model.users
	db := db.GetDB()
	UserID := userID
	db.Where("user_id = ?", UserID).First(&user)

	return user, nil
}
