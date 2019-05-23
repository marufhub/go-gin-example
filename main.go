package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/models"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/users")
	{
		v1.GET("/", models.GetUsers)
		v1.GET("/:user_id", models.GetUser)
	}
	r.Run()
}
