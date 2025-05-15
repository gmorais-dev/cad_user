package models

import (
	"time"
)

type Usuario struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nome      string    `gorm:"size:100" json:"nome"`
	Email     string    `gorm:"uniqueIndex;size:100" json:"email"`
	Senha     string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
