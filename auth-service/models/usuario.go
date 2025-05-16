package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `gorm:"size:100"`
	Email string `gorm:"uniqueIndex;size:100"`
	Senha string `json:"-"`
}
