package main

import (
	"campaign/auth"
	"campaign/config"
	"campaign/handler"
	"campaign/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDB()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	token, err := authService.ValidateToken("123eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.ZGs-2jfp3sxG2tFpsQkStWo7bH7zIr6ThpmjM4BpINA")
	if err != nil {
		fmt.Println("Error")
	}
	if token.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Error")
	}
	userHandler := handler.NewUserHandler(userService, authService)

	// r := routes.SetupRoutes(db)
	r := gin.Default()

	api := r.Group("/api/v1")
	api.POST("registration", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)
	api.POST("email_checkers", userHandler.CheckEmailAvailability)
	api.POST("avatars", userHandler.UploadAvatar)

	r.Run("localhost:9000")
}
