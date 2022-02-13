package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	// var jwtService service.JWTService = service.JWTAuthService()
	// var auth controllers.LoginController = controllers.LoginHandler(jwtService)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// r.POST("/login", func(ctx *gin.Context) {
	// 	auth.Login(ctx)
	// })

	// r.GET("/tasks", middleware.AuthorizeJWT(), controllers.FindTasks)
	// r.POST("/tasks", middleware.AuthorizeJWT(), controllers.CreateTask)
	// r.GET("/tasks/:id", middleware.AuthorizeJWT(), controllers.FindTask)
	// r.PATCH("/tasks/:id", middleware.AuthorizeJWT(), controllers.UpdateTask)
	// r.DELETE("tasks/:id", middleware.AuthorizeJWT(), controllers.DeleteTask)
	return r
}
