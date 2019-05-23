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

func init() {
	db := db.GetDB()
	db.AutoMigrate(&users{})
}

func getUsers(c *gin.Context) {
	var users []users
	db := db.GetDB()
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func getUser(c *gin.Context) {
	var user users
	user_id = c.Param("user_id")
	db.Where("user_id = ?", user_id).First(&user)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
