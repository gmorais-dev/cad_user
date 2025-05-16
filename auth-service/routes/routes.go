package routes

import (
	"auth-service/controllers"
	"auth-service/middlewares"
	"auth-service/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authService := services.NewAuthService()
	authController := controllers.NewAuthController(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetString("userID")
			c.JSON(200, gin.H{"user_id": userID, "message": "Protected profile route"})
		})
	}

	return r
}
