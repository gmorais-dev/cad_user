package services

import (
	"auth-service/dtos"
	"auth-service/models"
	"auth-service/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(dto *dtos.RegisterDTO) (*models.User, error) {
	// Verificar se usuário já existe
	var existingUser models.User
	if err := s.db.Where("email = ?", dto.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email já está em uso")
	}

	// Criar hash da senha
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	// Criar novo usuário
	user := models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("credenciais inválidas")
	}

	return &user, nil
}
