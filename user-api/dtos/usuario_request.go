package dtos

type UsuarioRequest struct {
	Nome  string `json:"nome" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"omitempty,min=6"`
}
