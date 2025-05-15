package controllers

import (
	"net/http"
	"strconv"
	"user-api/dtos"
	"user-api/helpers"
	"user-api/mappers"
	"user-api/services"

	"github.com/gin-gonic/gin"
)

func CriarUsuario(c *gin.Context) {
	var dto dtos.UsuarioRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		helpers.RespondError(c, http.StatusBadRequest, "Dados inválidos", err)
		return
	}

	usuario, err := services.CriarUsuario(dto)
	if err != nil {
		helpers.RespondError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	c.JSON(http.StatusCreated, mappers.MapUsuarioResponse(usuario))
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
		helpers.RespondError(c, http.StatusInternalServerError, "Erro ao listar usuários", err)
		return
	}

	var response []dtos.UsuarioResponse
	for _, usuario := range usuarios {
		response = append(response, mappers.MapUsuarioResponse(&usuario))
	}

	c.JSON(http.StatusOK, response)
}

func BuscarUsuarioPorID(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		helpers.RespondError(c, http.StatusBadRequest, "ID inválido", err)
		return
	}

	usuario, err := services.BuscarUsuarioPorID(id)
	if err != nil {
		helpers.RespondError(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, mappers.MapUsuarioResponse(usuario))
}

func AtualizarUsuario(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		helpers.RespondError(c, http.StatusBadRequest, "ID inválido", err)
		return
	}

	var dto dtos.UsuarioRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		helpers.RespondError(c, http.StatusBadRequest, "Dados inválidos", err)
		return
	}

	usuario, err := services.AtualizarUsuario(id, dto)
	if err != nil {
		helpers.RespondError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, mappers.MapUsuarioResponse(usuario))
}

func DeletarUsuario(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		helpers.RespondError(c, http.StatusBadRequest, "ID inválido", err)
		return
	}

	if err := services.DeletarUsuario(id); err != nil {
		helpers.RespondError(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.Status(http.StatusNoContent)
}
func parseIDParam(c *gin.Context) (uint64, error) {
	idStr := c.Param("id")
	return strconv.ParseUint(idStr, 10, 64)
}
