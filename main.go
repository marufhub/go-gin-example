package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/models"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./view/build", true)))
	v1 := r.Group("/api/users")
	{
		v1.GET("/", models.GetUsers)
		v1.GET("/:user_id", models.GetUser)
	}
	r.Run(":5000")
}
