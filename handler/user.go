package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marufhub/go-gin-example/models"
)

// Handle user request
// validate the request
// check basic authentication middleware
// call the usecase - function to data access
// transform the response to specific format using base handlers
// response error or successfull result.

// GetUsers return the list of all users
func GetUsers(c *gin.Context) {

	users := models.GetAllUsers()

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

//GetUser returns a user for the user_id
func GetUser(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	user := models.GetUserByID(userID)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
