package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marufhub/go-gin-example/middleware"
	"github.com/marufhub/go-gin-example/models"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./view/build", true)))
	api := r.Group("/api/users")
	{
		api.GET("/", models.GetUsers)
		api.GET("/:user_id", models.GetUser)
	}
	// basic authentication endpoints
	{
		basicAuth := api.Group("/admin")
		basicAuth.Use(middleware.AuthenticationRequired("admin"))
		{
			basicAuth.GET("/", models.GetUsers)
		}
	}
	r.Run(":5000")
}
