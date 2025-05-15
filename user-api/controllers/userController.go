package controllers

import (
	"net/http"
	"strconv"
	"user-api/dtos"
	"user-api/models"
	"user-api/services"

	"github.com/gin-gonic/gin"
)

// mapUsuarioResponse cria DTO de resposta a partir do modelo
func mapUsuarioResponse(usuario *models.Usuario) dtos.UsuarioResponse {
	return dtos.UsuarioResponse{
		ID:        uint64(usuario.ID),
		Nome:      usuario.Nome,
		Email:     usuario.Email,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
}

func CriarUsuario(c *gin.Context) {
	var dto dtos.UsuarioRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos", "detalhes": err.Error()})
		return
	}

	usuario, err := services.CriarUsuario(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapUsuarioResponse(usuario))
}

func ListarUsuarios(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	usuarios, err := services.ListarUsuarios(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var response []dtos.UsuarioResponse
	for _, usuario := range usuarios {
		response = append(response, mapUsuarioResponse(&usuario))
	}

	c.JSON(http.StatusOK, response)
}

func BuscarUsuarioPorID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	usuario, err := services.BuscarUsuarioPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapUsuarioResponse(usuario))
}

func AtualizarUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	var dto dtos.UsuarioRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos", "detalhes": err.Error()})
		return
	}

	usuario, err := services.AtualizarUsuario(id, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapUsuarioResponse(usuario))
}

func DeletarUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	if err := services.DeletarUsuario(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
