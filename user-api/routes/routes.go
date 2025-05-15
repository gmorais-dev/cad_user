package routes

import (
	"user-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuarioPorID)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)
}
