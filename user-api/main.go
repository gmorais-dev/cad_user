package main

import (
	"user-api/config"
	"user-api/controllers"
	"user-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConectaBanco()
	config.DB.AutoMigrate(&models.Usuario{})

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuarioPorID)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)

	r.Run(":8082")
}
