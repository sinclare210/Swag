package main

import (
	_ "github.com/swag/docs"
	"github.com/swag/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API
// @version 1
// @Description Sample description

// @contact.name Sinclare210
// @contact.url http://github.com/sinclare210
// @contact.email olajuwonsinclair@gmail.com

// @securityDefinitions.apiKey bearerToken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main(){
	r := gin.Default()

	v1 := r.Group("/api/v1")

	user := v1.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		user.POST("/",handlers.GetUsers)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

}
