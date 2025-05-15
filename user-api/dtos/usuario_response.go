package dtos

import "time"

type UsuarioResponse struct {
	ID        uint64    `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
