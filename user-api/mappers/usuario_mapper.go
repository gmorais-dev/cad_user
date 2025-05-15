package mappers

import (
	"user-api/dtos"
	"user-api/models"
)

func MapUsuarioResponse(usuario *models.Usuario) dtos.UsuarioResponse {
	if usuario == nil {
		return dtos.UsuarioResponse{}
	}

	return dtos.UsuarioResponse{
		ID:        uint64(usuario.ID),
		Nome:      usuario.Nome,
		Email:     usuario.Email,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
}
