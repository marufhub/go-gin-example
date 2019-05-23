package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type (
	users struct {
		gorm.Model
		UserID   int    `json:"user_id"`
		UserName string `json:"user_name"`
	}
)

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)

	}
	//Migrate the schema
	db.AutoMigrate(&users{})
}

func getUsers(c *gin.Context) {
	var users []users
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
func getUser(c *gin.Context) {
	var user users
	db.First(&user, 1)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
func main() {
	r := gin.Default()
	v1 := r.Group("/api/users")
	{
		v1.GET("/", getUsers)
		v1.GET("/1", getUser)
	}
	r.Run()
}
