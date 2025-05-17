package main

import (
	"log"

	"auth-service/config"
	"auth-service/controllers"
	"auth-service/models"
	"auth-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db := config.DB

	// MIGRAÇÃO DO MODELO
	//middlewares

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco: ", err)
	}

	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	router.Run(":8081")
}
