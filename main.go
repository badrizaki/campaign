package main

import (
	"campaign/config"
	"campaign/handler"
	"campaign/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDB()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// r := routes.SetupRoutes(db)
	r := gin.Default()

	api := r.Group("/api/v1")
	api.POST("registration", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)

	r.Run("localhost:9000")
}
