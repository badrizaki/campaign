package main

import (
	"campaign/auth"
	"campaign/campaign"
	"campaign/config"
	"campaign/handler"
	"campaign/helper"
	"campaign/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// r := routes.SetupRoutes(db)
	r := gin.Default()
	r.Static("/images", "./images")

	api := r.Group("/api/v1")
	api.POST("registration", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)
	api.POST("email_checkers", userHandler.CheckEmailAvailability)
	api.POST("avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("campaigns", campaignHandler.GetCampaigns)
	api.GET("campaign/:id", campaignHandler.GetCampaign)

	r.Run("localhost:9000")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
