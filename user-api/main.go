package main

import (
	"user-api/config"
	"user-api/controllers"
	"user-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta com o banco
	config.ConectaBanco()

	// Cria as tabelas automaticamente
	config.DB.AutoMigrate(&models.Usuario{})

	// Cria o servidor
	r := gin.Default()

	// Rota de criação de usuário
	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuarioPorID)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
