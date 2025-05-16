package main

import (
	"log"

	"auth-service/config"
	"auth-service/controllers"
	"auth-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta ao banco de dados
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Pega a instância do DB já conectada
	db := config.DB

	// Inicializa serviços
	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	// Configuração do router
	router := gin.Default()

	// Rotas de autenticação
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
