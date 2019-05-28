package usecase

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/db"
	"github.com/marufhub/go-gin-example/model"
)

// GetUsers return the list of all users
func GetUsers() ([]model.User, error) {
	var users []model.User
	db := db.GetDB()
	db.Find(&users)

	return users, nil

}

// GetUser returns a user for the user_id
func GetUser(userID string) (model.User, error) {
	var user model.User
	db := db.GetDB()
	db.Where("user_id = ?", userID).First(&user)

	return user, nil
}
