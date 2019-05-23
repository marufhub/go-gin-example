package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/db"
)

type (
	users struct {
		gorm.Model
		UserID   int    `json:"user_id"`
		UserName string `json:"user_name"`
	}
)

// GetUsers return the list of all users
func GetUsers(c *gin.Context) {
	var users []users
	db := db.GetDB()
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// GetUser returns a user for the user_id
func GetUser(c *gin.Context) {
	var user users
	db := db.GetDB()
	UserID := c.Param("user_id")
	db.Where("user_id = ?", UserID).First(&user)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
