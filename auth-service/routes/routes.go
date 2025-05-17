package routes

import (
	"auth-service/config"
	"auth-service/controllers"
	"auth-service/middlewares"
	"auth-service/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Inicializa serviço e controller
	authService := services.NewAuthService(config.DB)
	authController := controllers.NewAuthController(authService)

	// Rotas públicas
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Rotas protegidas com JWT
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, exists := c.Get("userID")
			if !exists {
				c.JSON(401, gin.H{"error": "usuário não autenticado"})
				return
			}

			c.JSON(200, gin.H{
				"user_id": userID,
				"message": "Protected profile route",
			})
		})
	}

	return r
}
